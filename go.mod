module github.com/docker/cli

go 1.25.0

require (
	dario.cat/mergo v1.0.2
	github.com/containerd/errdefs v1.0.0
	github.com/containerd/log v0.1.0
	github.com/containerd/platforms v1.0.0-rc.4
	github.com/creack/pty v1.1.24
	github.com/distribution/reference v0.6.0
	github.com/docker/cli-docs-tool v0.11.0
	github.com/docker/distribution v2.8.3+incompatible
	github.com/docker/docker-credential-helpers v0.9.8
	github.com/docker/go-connections v0.7.0
	github.com/docker/go-units v0.5.0
	github.com/fvbommel/sortorder v1.1.0
	github.com/go-jose/go-jose/v4 v4.1.4
	github.com/go-viper/mapstructure/v2 v2.5.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.7.0
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/google/uuid v1.6.0
	github.com/mattn/go-runewidth v0.0.24
	github.com/moby/go-archive v0.2.0
	github.com/moby/moby/api v1.55.0-rc.1
	github.com/moby/moby/client v0.5.0-rc.1
	github.com/moby/patternmatcher v0.6.1
	github.com/moby/swarmkit/v2 v2.1.2
	github.com/moby/sys/atomicwriter v0.1.0
	github.com/moby/sys/capability v0.4.0
	github.com/moby/sys/sequential v0.7.0
	github.com/moby/sys/signal v0.7.1
	github.com/moby/sys/symlink v0.3.0
	github.com/moby/term v0.5.2
	github.com/morikuni/aec v1.1.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.1.1
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
	github.com/sirupsen/logrus v1.9.4
	github.com/spf13/cobra v1.10.2
	github.com/spf13/pflag v1.0.10
	github.com/tonistiigi/go-rosetta v0.0.0-20220804170347-3f4430f2d346
	github.com/xeipuuv/gojsonschema v1.2.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.69.0
	go.opentelemetry.io/otel v1.44.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.44.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.44.0
	go.opentelemetry.io/otel/metric v1.44.0
	go.opentelemetry.io/otel/sdk v1.44.0
	go.opentelemetry.io/otel/sdk/metric v1.44.0
	go.opentelemetry.io/otel/trace v1.44.0
	go.yaml.in/yaml/v3 v3.0.4
	golang.org/x/sync v0.21.0
	golang.org/x/sys v0.46.0
	golang.org/x/term v0.44.0
	golang.org/x/text v0.38.0
	gotest.tools/v3 v3.5.2
	tags.cncf.io/container-device-interface v1.1.0
)

// Use our WASM-patched fork of moby/term.
replace github.com/moby/term => github.com/justwasm/mobyterm v0.5.2-wasm

// Resolve google.golang.org/genproto ambiguous import between
// the old monolithic module and the new sub-modules.
replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
