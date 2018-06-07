# Download SRA

A simple download tool to download `.sra` file from a repository of [INSDC](http://insdc.org) members.

## Why not `prefetch`?

Because we are living far away from Maryland. `download_sra` can choose a repository to be used.

## Usage

```
$ download_sra SRR000001
$ download_sra SRR000001 NCBI
$ download_sra SRR000001 EBI
$ download_sra SRR000001 DDBJ
```
