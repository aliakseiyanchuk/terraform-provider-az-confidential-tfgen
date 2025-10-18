# `general` group

This is a container group for data sources and resources that are not specific to a particular Azure service.
The information about this group can be obtained using the following command:

```shell
tfgen general help
```

The command above would print:
```
Usage: tfgen [<standard options>] general <subcommand> [<args>]
Possible sub-commands are:
- content
```

Currently the group contains a single resource, [content datasource](content.md). 