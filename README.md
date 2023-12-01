# Prometheus Data Mock Tool
This tool is designed to mock Prometheus data using configuration files. It allows you to generate synthetic data for testing and development purposes.

## Configuration
The tool's configuration is defined in a YAML file. The following sections describe the available options:

``` yaml
template_dir: ./templates
output_dir: ./output
template_value_dir: ./template_values

global_config:
  template_value_path: ./template_values_without_unix.json
  days: 1
  resolution_seconds: 15
  endtime_unix: 1698048047

template_without_unix_Configs:
  - name: nvidia_gpu_exporter
    template_value_path: gpu_node_2.json

template_with_unix_configs:
  - name: nvidia_gpu_exporter_with_unix
    template_value_path: samples_gpu_node_1.json
```

### Global Configuration
The global configuration section contains settings that apply to all templates. These settings are specified under the global_config key. The available options are:

* template_value_path: The path to the JSON file containing template values. This file provides the data used to generate the mock data. By default, the tool looks for this file in the template_values directory.
* days: The number of days for which data will be generated. Default value is 30.
* resolution_seconds: The resolution, in seconds, of the generated data. Default value is 15.
* endtime_unix: The Unix timestamp representing the end time of the generated data. Default value is 1698048047.

### Template Configuration
The template configurations define individual templates and their specific settings. There are two types of template configurations:

#### Templates without Unix timestamp
Templates without Unix timestamp are defined under the template_without_unix_configs key. Each template configuration consists of the following options:

* name: The name of the template.
* template_value_path: The path to the JSON file containing template values specific to this template. By default, the tool looks for this file in the template_values directory.

#### Templates with Unix timestamp
* Templates with Unix timestamp are defined under the template_with_unix_configs key. Each template configuration consists of the following options:
* name: The name of the template.
* template_value_path: The path to the JSON file containing template values specific to this template. By default, the tool looks for this file in the template_values directory.

## Build

``` shell
make build
```

## Usage
Ensure that you have the necessary configuration files in the appropriate directories:

Run the tool using the following command:

``` shell
./prometheus-data-generator 
Dec  1 03:46:22.488 nvidia_gpu_exporter, 2023-10-22T08:00:47Z -> 2023-10-23T08:00:47Z, step: 15s
Dec  1 03:46:22.488 progress: 0%, 2023-10-22T08:00:47Z -> 2023-10-22T10:00:00Z
Dec  1 03:46:22.488 process template...
Dec  1 03:46:22.722 create blocks...
BLOCK ULID                  MIN TIME                       MAX TIME                       DURATION     NUM SAMPLES  NUM CHUNKS   NUM SERIES   SIZE
01HGHQR4GFSFXM8Q5SSNW7Y1RJ  2023-10-22 08:00:47 +0000 UTC  2023-10-22 09:59:47 +0000 UTC  1h59m0.001s  30528        256          64           72KiB840B
Dec  1 03:46:22.950 progress: 8%, 2023-10-22T10:00:00Z -> 2023-10-22T12:00:00Z
Dec  1 03:46:22.950 process template...
Dec  1 03:46:23.118 create blocks...
```

The tool will generate the mock data based on the provided configuration and save it in the output directory.

That's it! You can now use this tool to mock Prometheus data for your testing and development needs.