# Download SRA

[![DOI](https://zenodo.org/badge/136470994.svg)](https://zenodo.org/badge/latestdoi/136470994)

A simple download tool to download `.sra` file from a repository of [INSDC](http://insdc.org) members.

## Why not `prefetch`?

Because we are living far away from Maryland. `download-sra` can choose a repository to be used.

## Prerequisites

- `curl`
- [`prefetch`](https://github.com/ncbi/sra-tools/wiki/01.-Downloading-SRA-Toolkit) (to download data from NCBI)

## Usage

```
$ ./download-sra
usage: download_sra [-r|--repo] [ncbi|ebi|ddbj] <SRA Run ID>[ <Run ID>..]
```

Try with `SRR1274307` which is a .sra file less than 1MB.
Get data from DDBJ:

```
download-sra SRR1274307
```

or

```
download-sra -r DDBJ SRR1274307
```

Use EBI:

```
download-sra -r EBI SRR1274307
```

## Docker Container

Available on [GHCR](https://github.com/inutano/download-sra/pkgs/container/download-sra)

```
docker run --rm -it -v $(pwd):/work -w /work docker pull ghcr.io/inutano/download-sra:cb2bba4 download-sra -r "DDBJ" "SRR1274307"
```

```
docker run --rm -it -v $(pwd):/work -w /work docker pull ghcr.io/inutano/download-sra:cb2bba4 download-sra -r "EBI" "SRR1274307"
```
