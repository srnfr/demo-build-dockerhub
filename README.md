# demo-build-dockerhub

Dumb demo to demonstrate automatic build between 
- DockerHub
- GitHub Actions & Buildjet with push to [ghcr.io/srnfr/demo-build-dockerhub:latest](https://ghcr.io/srnfr/demo-build-dockerhub:latest)

Cosign signature added.


Pour vérifier :
- la signature de l'image : ```cosign verify --key cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest```
- le SBOM : ```cosign download sbom ghcr.io/srnfr/demo-build-dockerhub --output-file monsbom.sbom```
