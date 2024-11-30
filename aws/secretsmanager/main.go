package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		slog.Error(fmt.Errorf("get user home dir: %w", err).Error())
		os.Exit(1)
	}

	profiles := strings.Split(os.Getenv("AWS_PROFILES"), ",")
	for _, profile := range profiles {
		err = writeEnvFile(strings.TrimSpace(profile), dirname)
		if err != nil {
			slog.Error(fmt.Errorf("write env file: %w", err).Error())
			os.Exit(1)
		}
	}
}

func writeEnvFile(profile, prefix string) error {
	slog.Info("write env file", slog.String("profile", profile))

	client, err := NewSecretsClient(profile)
	if err != nil {
		return fmt.Errorf("create secret grabber: %w", err)
	}

	secrets, err := client.getSecrets()
	if err != nil {
		return fmt.Errorf("grab secrets: %w", err)
	}

	data := []byte(strings.Join(secrets[:], "\n") + "\n")

	err = os.WriteFile(prefix+"/"+profile+".env", data, 0644)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

type SecretsClient struct {
	profile             string
	currentVersionStage *string

	secretsManager *secretsmanager.SecretsManager
}

func NewSecretsClient(profile string) (*SecretsClient, error) {
	c := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION"))

	s, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	})
	if err != nil {
		return nil, fmt.Errorf("new session: %w", err)
	}

	_, err = sts.New(s).GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, fmt.Errorf("get caller identity: %w", err)
	}

	return &SecretsClient{
		profile:             profile,
		currentVersionStage: aws.String("AWSCURRENT"),
		secretsManager:      secretsmanager.New(s, c),
	}, nil
}

func (client *SecretsClient) getSecrets() ([]string, error) {
	newline := regexp.MustCompile(`\r?\n`)

	names, err := client.getSecretNames()
	if err != nil {
		return nil, fmt.Errorf("grab secret names: %w", err)
	}
	slices.Sort(names)

	var secrets []string
	secrets = append(secrets, "# "+client.profile)

	for _, name := range names {
		secrets = append(secrets, "\n# "+name)

		input := &secretsmanager.GetSecretValueInput{
			SecretId:     &name,
			VersionStage: aws.String("AWSCURRENT"),
		}
		value, err := client.secretsManager.GetSecretValue(input)
		if err != nil {
			slog.Debug(fmt.Errorf("get secret value: %w", err).Error())
			continue
		}

		var s string
		if value.SecretString != nil {
			s = *value.SecretString
		} else {
			continue
		}

		var m map[string]string
		err = json.Unmarshal([]byte(s), &m)
		if err != nil {
			slog.Debug(fmt.Errorf("unmarshall secret json string: %w", err).Error())
			continue
		}

		for k, v := range m {
			secrets = append(secrets, fmt.Sprintf(`%s="%s"`, k, newline.ReplaceAllString(v, "\\n")))
		}
	}

	return secrets, nil
}

func (client *SecretsClient) getSecretNames() ([]string, error) {
	var names []string

	arr, nextToken, err := client.getSecretNamesWithToken(nil)
	if err != nil {
		return nil, fmt.Errorf("grab secret names with token: %w", err)
	}
	names = append(names, arr...)

	for nextToken != nil && len(arr) > 0 {
		arr, nextToken, err = client.getSecretNamesWithToken(nextToken)
		if err != nil {
			return nil, fmt.Errorf("grab secret names with token: %w", err)
		}
		names = append(names, arr...)
	}

	return names, nil
}

func (client *SecretsClient) getSecretNamesWithToken(nextToken *string) ([]string, *string, error) {
	input := secretsmanager.ListSecretsInput{
		NextToken: nextToken,
	}
	secrets, err := client.secretsManager.ListSecrets(&input)
	if err != nil {
		return nil, nil, fmt.Errorf("list secrets: %w", err)
	}

	var names []string
	for _, secret := range secrets.SecretList {
		names = append(names, *secret.Name)
	}

	return names, secrets.NextToken, nil
}
