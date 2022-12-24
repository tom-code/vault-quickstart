
vault auth enable kubernetes
vault write auth/kubernetes/config kubernetes_host=https://kubernetes.default:443

vault write auth/kubernetes/role/role1 \
    bound_service_account_names=\* \
    bound_service_account_namespaces=\* \
    policies=default \

vault policy read default > /tmp/a.hcl

cat << EOF >> /tmp/a.hcl
path "secret/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}
EOF

vault policy write default /tmp/a.hcl 


vault kv put -mount=secret test1 foo=world

