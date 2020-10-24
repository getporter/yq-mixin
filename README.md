# yq Mixin for Porter

This is a mixin for Porter that provides the [yq] CLI.

[![Build Status](https://dev.azure.com/getporter/porter/_apis/build/status/yq-mixin?branchName=main)](TODO)

## Mixin Configuration

When you declare the mixin, you can also configure additional extensions to install

**Specify the yq client version**
```yaml
mixins:
- yq:
    clientVersion: "3.4.0"
```

## Mixin Syntax and Examples

See the [yq documentation](https://mikefarah.gitbook.io/yq/) for the supported commands.

All commands support [suppress-output](#suppress-output) and [outputs].

### Suppress Output

The `suppress-output` field controls whether output from the mixin should be
prevented from printing to the console. By default this value is false, using
Porter's default behavior of hiding known sensitive values. When 
`suppress-output: true` all output from the mixin (stderr and stdout) are hidden.

Step outputs (below) are still collected when output is suppressed. This allows
you to prevent sensitive data from being exposed while still collecting it from
a command and using it in your bundle.

### Outputs

The mixin supports `jsonpath` and `path` outputs.


#### JSON Path

The `jsonPath` output treats stdout like a json document and applies the expression, saving the result to the output.

```yaml
outputs:
- name: NAME
  jsonPath: JSONPATH
```

For example, if the `jsonPath` expression was `$[*].id` and the command sent the following to stdout: 

```json
[
  {
    "id": "1085517466897181794",
    "name": "my-vm"
  }
]
```

Then then output would have the following contents:

```json
["1085517466897181794"]
```

#### File Paths

The `path` output saves the content of the specified file path to an output.

```yaml
outputs:
- name: kubeconfig
  path: /root/.kube/config
```

### Command

Run any yq command supported by the CLI.

#### Command Syntax

```yaml
yq:
  description: "Description of the command"
  arguments:
  - ARG1
  - ARG2
  flags:
    FLAGNAME: FLAGVALUE
    REPEATED_FLAG:
    - FLAG_VALUE1
    - FLAG_VALUE2
  suffix-arguments:
  - SUFFIX_ARG1
  suppress-output: BOOL # Defaults to false
  outputs:
    - name: NAME
      jsonPath: JSONPATH
    - name: NAME
      path: SOURCE_FILEPATH
```

#### Command Example

Change the value of b.c to cat and then capture the output.

**config.yaml**
```yaml
b:
  c: dog
```

```yaml
yq:
  description: "Set value and remember the output"
  arguments:
    - write
    - config.yaml
  flags:
    indent: 4
  suffix-arguments:
    - b.c
    - cat
  outputs: # Capture all of stdout for use in a later step
    - name: vars
      regex: "(.)"
```

### Update

This is a convenience function that simplifies a common task: reading in a file
and editing it in-place. More advanced scenarios should use the write command.

The update command applies the expression to all documents defined in the file
and edits the file in-place.

#### Update Syntax

```yaml
yq:
  description: "Description of the command"
  update:
    file: PATH OF FILE TO UPDATE
    expression: YQ PATH EXPRESSION
    value: REPLACEMENT VALUE
    style: STYLE # formatting style of the value: single, double, folded, flow, literal, tagged
  suppress-output: BOOL # Defaults to false
  outputs:
    - name: NAME
      jsonPath: JSONPATH
    - name: NAME
      path: SOURCE_FILEPATH
```

* [Path expression syntax](https://mikefarah.gitbook.io/yq/usage/path-expressions)

#### Update Example

Update the image in a Kubernetes deployment with the image inside the bundle.

```yaml
yq:
  description: "Replace deployment image"
  update:
    file: manifests/deployment.yaml
    expression: spec.template.spec.containers.(name==webserver).image
    value: "{{bundle.images.webserver.repository}}@{{bundle.images.webserver.digest}}"
```

### Write a value

Wrapper for the `yq write` command.

```yaml
yq:
  description: "Description of the command"
  write:
    file: FILEPATH # Input file
    expression: YQ PATH EXPRESSION
    value: REPLACEMENT VALUE
    style: STYLE # formatting style of the value: single, double, folded, flow, literal, tagged
    script: PATH OF SCRIPT # yaml script for updating yaml
    document: INTEGER # Defaults to 0
    inplace: BOOL # Defaults to false
    destination: FILEPATH # Saves stdout to the specified file, used when inplace is false
  prettyPrint: BOOL # Defaults to false
  toJson: BOOL # Defaults to false
  suppress-output: BOOL  # Defaults to false
  outputs:
    - name: NAME
      jsonPath: JSONPATH
    - name: NAME
      path: SOURCE_FILEPATH
```

```yaml
yq:
  description: "Replace deployment image"
  write:
    file: manifests/deployment.yaml
    document: 1 # Modify the second document in the file. Defaults to 0
    expression: spec.template.spec.containers.(name==webserver).image
    value: "{{bundle.images.webserver.repository}}@{{bundle.images.webserver.digest}}"
    destination: final-deployment.yaml
```

### Merge files

Wrapper for the `yq merge` command.

#### Merge Syntax

```yaml
yq:
  description: "Description of Command"
  merge:
    files:
      - FILEPATH_1
      - FILEPATH_2
    append: BOOL # Defaults to false
    autocreate: BOOL # Defaults to true
    inplace: BOOL # Defaults to false
    overwrite: BOOL # Defaults to false
    destination: FILEPATH # Saves stdout to the specified file, used when inplace is false
  prettyPrint: BOOL # Defaults to false
  toJson: BOOL # Defaults to false
  suppress-output: BOOL  # Defaults to false
  outputs:
    - name: NAME
      jsonPath: JSONPATH
    - name: NAME
      path: SOURCE_FILEPATH  
```

#### Merge Example

Merge two YAML files and save the result to a json file.

```yaml
yq:
  description: "Merge environment configuration"
  merge:
    files:
      - base.yaml
      - staging.yaml
    destination: final.json
  toJson: true
  prettyPrint: true
```

[yq]: https://github.com/mikefarah/yq/releases
