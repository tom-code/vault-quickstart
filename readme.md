
# vault k8s developer quick start

- install vault from helm chart https://github.com/hashicorp/vault-helm
    - in values.yaml enable dev mode


- enable k8s authentication, add role, modify acl and create secret
    - execute steps in vault.sh inside vault pod

- now any pod can access any secret - see app.go