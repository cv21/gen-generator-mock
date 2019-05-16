# gen-generator-mock

It is a generator for [gen](https://github.com/cv21/gen) which is able to generate mocks.

Generates [mockery](https://github.com/vektra/mockery)-compatible code for [testify](https://github.com/stretchr/testify).

Parameters:

| Name | Description |
|----------|----------|
| interface_name      | The name of the interface to be mocked      |
| out_path      | Output path for the generated file      |
| package_name     | Generated file package name     |
| mock_struct_name_template     | Template of mock structure     |


Config example:

```json
{
  "files": [
    {
      "path": "./service.go",
      "generators": [
        {
          "name": "mock",
          "version": "1.0.0",
          "params": {
            "interface_name": "StringService",
            "out_path_template": "./generated/%s_mock_gen.go",
            "source_package_path": "github.com/cv21/gen/examples/stringsvc",
            "target_package_path": "github.com/cv21/gen/examples/stringsvc/generated",
            "mock_struct_name_template": "%sMock"
          }
        }
      ]
    }
  ]
}
```

See examples directory to discover generated code.