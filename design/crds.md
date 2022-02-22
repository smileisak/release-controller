```yaml
---
title: Release Controller CRDs
version: v1alpha1
authors: Ismail KABOUBI
creation-date: 2020-09-01
status: draft
---
```

# Release Controller CRDs

Release Controller is a Kubernetes Controller that tracks new releases from multiple sources and perform an action
when a new release is out. This controller can be useful for example to track upstream releases and create actions
to be up-to-date regarding what is happening in open source field.

## Proposal

### API

### ReleaseTrack

`ReleaseTrack` Custom Resource Definition is cluster scoped, it defines the following :
1. releaseStoreRef from where the releases will be grabbed.
2. action defines what to do when a new Release is out.


```yaml
apiVersion: release-controller.io/v1alpha1
kind: ReleaseTrack 
metadata:
  name: argocd
spec:
  releaseStoreRef:
    name: github
  action:
    provider:
      slack:
        auth:
          secretRef:
            name: secret-containing-token
            namespace: example
status:
  lastVersion: v2.1.1
```

Action Providers can be multiple as following:

* Slack
* Teams
* Gitlab Issue
* Jira Ticket

### ReleaseStore

```yaml
apiVersion: release-controller.io/v1alpha1
kind: ReleaseStore 
metadata:
  name: github
spec:
  provider:
    github:
      repository: argoproj/argo-cd
      tagRegex: "v\d+.\d+.\d+$"
status:
  todo: todo
```

Providers can be multiple, but ReleaseStore should refer to one provider. These providers are relevant for the project:

* Github
* Gitlab
* ArtifactHub



ActionStores can be multiple as following:

* Slack
* Teams
* Gitlab Issue
* Jira Ticket