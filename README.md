# demo-build-dockerhub

Dumb demo to demonstrate automatic build with 
- Go lang compilatation & test on every related push
- Docker image building (including go compilation inside Dockerfile) on every new tag
- CI with GitHub Actions & Buildjet with push to [ghcr.io/srnfr/demo-build-dockerhub:latest](https://ghcr.io/srnfr/demo-build-dockerhub:latest)
- Asynchronous build with DockerHub registry triggered by Web Hook

[Cosign](https://docs.sigstore.dev/cosign/installation/#container-images) signature added.

SBOM attachment added.

In order to   :

a) check the image's signature : 

```cosign verify --key cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest```

b) download the SBOM : 

```cosign download sbom ghcr.io/srnfr/demo-build-dockerhub --output-file monsbom.sbom```

Then, analyze the vulns with [grype](https://github.com/anchore/grype) : 
```grype monsbom.sbom```

__Tip__: use the cosign Docker image  (rather than the complex install) :

a) Verify

```bash
docker run gcr.io/projectsigstore/cosign verify --key https://raw.githubusercontent.com/srnfr/demo-build-dockerhub/master/cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest
```

b) Download the SBOM

```bash
docker run -v cosign_data:/home/nonroot gcr.io/projectsigstore/cosign download sbom ghcr.io/srnfr/demo-build-dockerhub --output-file monsbom.sbom
```

c) Locate the ```monsbom.sbom``` on the Host's system :

```bash
docker volume inspect cosign_data
```

(usually ```/var/lib/docker/volumes/cosign_data/_data```)

d) Install grype 
```bash
curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /usr/local/bin
```

e) Run 
```bash
grype monsbom.sbom
```
