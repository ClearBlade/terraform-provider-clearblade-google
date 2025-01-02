package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// https://raw.githubusercontent.com/ClearBlade/helm-charts/refs/tags/clearblade-iot-enterprise-3.0.4/gke-default-values.yaml
const (
	helmTemplateURLPrefix = "https://raw.githubusercontent.com/ClearBlade/helm-charts/refs/tags/"
	helpTemplateURLSuffix = "/gke-default-values.yaml"
)

var (
	_ function.Function = GetHelmYamlTemplate{}
)

func NewGetHelmYamlTemplateFunction() function.Function {
	return GetHelmYamlTemplate{}
}

type GetHelmYamlTemplate struct{}

func (g GetHelmYamlTemplate) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "get_helm_yaml_template"
}

func (g GetHelmYamlTemplate) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Function to get a Helm YAML template file",
		MarkdownDescription: "Gets a Helm YAML template from a GitHub release",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "helm_chart",
				MarkdownDescription: "Location of the Helm chart on Github",
			},
		},
		Return: function.StringReturn{},
	}
}

func (g GetHelmYamlTemplate) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var helmChartURL string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &helmChartURL))
	if resp.Error != nil {
		return
	}

	release, err := getRelease(helmChartURL)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}
	file, err := getTemplateFile(release)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Result.Set(ctx, string(file))
}

func getTemplateFile(release string) ([]byte, error) {
	url := helmTemplateURLPrefix + release + helpTemplateURLSuffix
	c := http.Client{Timeout: time.Duration(60) * time.Second}
	resp, err := c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read HTTP response: %w", err)
	}
	return body, nil
}

func getRelease(url string) (string, error) {
	split := strings.Split(url, "/")
	if len(split) < 9 {
		return "", fmt.Errorf("Invalid helm chart url: %s", url)
	}
	return split[7], nil
}
