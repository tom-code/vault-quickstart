
package main

import (
	"context"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/kubernetes"
)

func stringlistFromSecrets(secret *vault.Secret) []string {
	out := []string{}
	keys, ok := secret.Data["keys"]
	if ok {
		log.Println("a1")
		keysarr, ok := keys.([]interface{})
		if ok {
			for _, keyi := range(keysarr) {
				switch keyi.(type) {
				case string:
					out = append(out, keyi.(string))
				}
			}
		}
	}
	return out
}

func test(url string, role string) error {
	config := vault.DefaultConfig()
	config.Address = url

	client, err := vault.NewClient(config)
	if err != nil {
		return fmt.Errorf("[vault] unable to initialize Vault client: %w", err)
	}

	k8sAuth, err := auth.NewKubernetesAuth(
		role,
		auth.WithServiceAccountTokenPath("/var/run/secrets/kubernetes.io/serviceaccount/token"),
	)
	if err != nil {
		return fmt.Errorf("[vault] unable to initialize Kubernetes auth method: %w", err)
	}

	authInfo, err := client.Auth().Login(context.TODO(), k8sAuth)
	if err != nil {
		return fmt.Errorf("[vault] unable to log in with Kubernetes auth: %w", err)
	}
	if authInfo == nil {
		return fmt.Errorf("[vault] no auth info was returned after login")
	}

	s, err := client.Logical().List("secret/metadata")
	secrets := stringlistFromSecrets(s)

	log.Printf("secrets: %v", secrets)

	secret, err := client.KVv2("secret").Get(context.Background(), "test1")
	if err != nil {
		return fmt.Errorf("[vault] unable to read secret: %w", err)
	}
	log.Println(secret)
	return nil
}

func main() {
	err := test("http://vault.vault:8200", "role1")
	log.Println(err)
}
