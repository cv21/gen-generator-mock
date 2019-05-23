# gen-generator-mock


#### Generator Parameters
| Parameter Key | Example | Description |
| --- | --- | --- |
|interface_name|SomeService|This is name of interface to mock.|
|out_path_template|./generated/%s_mock_gen.go|It is output path template. It applies %s literal which holds interface name in snack_case.|
|source_package_path|github.com/cv21/gen-generator-mock/examples/stringsvc|It is package of source file.|
|target_package_path|github.com/cv21/gen-generator-mock/examples/stringsvc/bla|It is target package path.|
|mock_struct_name_template|MyPrettyMockOf%s|It is a template for custom struct naming. It applies %s literal which holds interface name.|

#### Config Example:

```json
{
    "files": [
        {
            "path": "...",
            "generators": [
                {
                    "repository": "github.com/cv21/gen-generator-mock",
                    "version": "v1.0.0",
                    "params": {
                        "interface_name":"SomeService",
                        "out_path_template":"./generated/%s_mock_gen.go",
                        "source_package_path":"github.com/cv21/gen-generator-mock/examples/stringsvc",
                        "target_package_path":"github.com/cv21/gen-generator-mock/examples/stringsvc/bla",
                        "mock_struct_name_template":"MyPrettyMockOf%s"
                    }
                }
            ]
        }
    ]
}
```

