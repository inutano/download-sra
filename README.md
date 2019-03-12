# Download SRA

A simple download tool to download `.sra` file from a repository of [INSDC](http://insdc.org) members.

## Why not `prefetch`?

Because we are living far away from Maryland. `download_sra` can choose a repository to be used.

## Prerequisites

- `wget`
- `curl`
- [`jq`](https://stedolan.github.io/jq/)

## Usage

```
$ download_sra SRR000001
$ download_sra SRR000001 NCBI
$ download_sra SRR000001 EBI
$ download_sra SRR000001 DDBJ
```

## Docker Container

Available on [Quay.io](https://quay.io/repository/inutano/download-sra)

```
$ docker run --rm -it -v $(pwd):/work -w /work quay.io/inutano/download-sra:0.1.2 "DRR000001" "NCBI"
$ docker run --rm -it -v $(pwd):/work -w /work quay.io/inutano/download-sra:0.1.2 "DRR000001" "EBI"
$ docker run --rm -it -v $(pwd):/work -w /work quay.io/inutano/download-sra:0.1.2 "DRR000001" "DDBJ"
```
