# demo-build-dockerhub

Dumb demo to demonstrate automatic build between 
- DockerHub
- GitHub Actions & Buildjet with push to [ghcr.io/srnfr/demo-build-dockerhub:latest](https://ghcr.io/srnfr/demo-build-dockerhub:latest)

[Cosign](https://docs.sigstore.dev/cosign/installation/#container-images) signature added.
SBOM attachment added.


Pour vérifier :
- la signature de l'image : ```cosign verify --key cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest```
- le SBOM : ```cosign download sbom ghcr.io/srnfr/demo-build-dockerhub --output-file monsbom.sbom```
- Pour analyser les vulnérabilités avec [grype](https://github.com/anchore/grype) : ```grype monsbom.sbom```

PS: pour tester en utilisant l'image Docker cosign (+ simple que faire l'installation) :

```docker run gcr.io/projectsigstore/cosign verify --key https://raw.githubusercontent.com/srnfr/demo-build-dockerhub/master/cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest```
