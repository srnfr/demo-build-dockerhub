# demo-build-dockerhub

Dumb demo to demonstrate automatic build with 
- GOlang build inside the image
- DockerHub registry
- GitHub Actions & Buildjet with push to [ghcr.io/srnfr/demo-build-dockerhub:latest](https://ghcr.io/srnfr/demo-build-dockerhub:latest)

[Cosign](https://docs.sigstore.dev/cosign/installation/#container-images) signature added.

SBOM attachment added.


In order to   :

a) check the image's signature : 
```cosign verify --key cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest```

b) download the SBOM : 
```cosign download sbom ghcr.io/srnfr/demo-build-dockerhub --output-file monsbom.sbom```

Then, analyse the vulns with [grype](https://github.com/anchore/grype) : 
```grype monsbom.sbom```

__Tip__: use the cosign Docker image  (rather than the complex install) :

a) Verify
```docker run gcr.io/projectsigstore/cosign verify --key https://raw.githubusercontent.com/srnfr/demo-build-dockerhub/master/cosign.pub ghcr.io/srnfr/demo-build-dockerhub:latest```

b) Download the SBOM
```docker run -v cosign_data:/home/nonroot gcr.io/projectsigstore/cosign download sbom ghcr.io/srnfr/demo-build-dockerhub --output-file monsbom.sbom```

c) Locate the ```monsbom.sbom``` on the Host's system :
```docker volume inspect cosign_data```

d) Run ```grype monsbom.sbom```
