# Setting up Github Actions

Setting up the github actions to work with Azure.

In order to auth into the account we needed to make a service principal.

```bash
az ad sp create-for-rbac --name HomesDeploymentPrincipal --role Contributor --scopes /subscriptions/{subscription_id}
```

For me I had to fetch the subscription id of my given subscription by running

```bash
az account list 
```

and looking for the subscription that is useful.

On the Azure UI I had to create a ResourceGroup and create a container registry. 
