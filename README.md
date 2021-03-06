ipfs-ds-convert
==================

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![Coverage Status](https://coveralls.io/repos/github/ipfs/ipfs-ds-convert/badge.svg)](https://coveralls.io/github/ipfs/ipfs-ds-convert)
[![Travis CI](https://circleci.com/gh/ipfs/ipfs-ds-convert/tree/master.svg?style=shield)](https://circleci.com/gh/ipfs/ipfs-ds-convert/tree/master)

> Datastore converter for go-ipfs

This tool is WIP and may damage your data. Make sure to make a backup first.

TODO:
- [x] Finish basic conversion code
- [x] package.json for gx
- [ ] Tests
  - [x] CI (needs https://github.com/ipfs/go-ipfs/pull/4007, https://github.com/ipfs/go-ipfs/pull/3575)
  - [ ] Coverage > 80% or more
- [ ] Review
- [ ] Standard readme
- [x] Revert on error / from backup
- [x] Cleanup backup subcommand
- [x] Optimize some standard cases
  - [x] Don't copy directories when not needed
  - [ ] Detect renames
    - Not that common
- [x] Report progress
- [ ] Don't depend on go-ipfs

## Contribute

PRs are welcome!

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Łukasz Magiera
