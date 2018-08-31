package plugins

// GCSOutput CRD name
const GCSOutput = "gcs"

// GCSDefaultValues for Amazaon S3 output plugin
var GCSDefaultValues = map[string]string{
	"bufferTimeKey":  "3600",
	"bufferTimeWait": "10m",
	"bufferPath":     "/buffers/gcs",
	"format":         "json",
}

// GCSTemplate for Amazaon S3 output plugin
const GCSTemplate = `
<match {{ .pattern }}.**>
  @type gcs

  project {{ .project }}
  credentialsJson {{ .credentials }}
  bucket {{ .bucket }}
  object_key_format %{path}%{time_slice}_%{index}.%{file_extension}
  path logs/${tag}/%Y/%m/%d/

  # if you want to use ${tag} or %Y/%m/%d/ like syntax in path / object_key_format,
  # need to specify tag for ${tag} and time for %Y/%m/%d in <buffer> argument.
  <buffer tag,time>
    @type file
    path /buffers/gcs
    timekey 1h # 1 hour partition
    timekey_wait 10m
    timekey_use_utc true # use utc
  </buffer>

  <format>
    @type json
  </format>
</match>`
