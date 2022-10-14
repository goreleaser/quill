package commands

import (
	"debug/macho"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/anchore/quill/cmd/quill/cli/application"
	"github.com/anchore/quill/cmd/quill/cli/options"
	"github.com/anchore/quill/quill"
)

var _ options.Interface = &signConfig{}

type signConfig struct {
	Path            string `yaml:"path" json:"path" mapstructure:"path"`
	options.Signing `yaml:"signing" json:"signing" mapstructure:"signing"`
}

func Sign(app *application.Application) *cobra.Command {
	opts := &signConfig{
		Signing: options.DefaultSigning(),
	}

	cmd := &cobra.Command{
		Use:   "sign PATH",
		Short: "sign a macho (darwin) executable binary",
		Example: options.FormatPositionalArgsHelp(
			map[string]string{
				"PATH": "the darwin binary to sign",
			},
		),
		Args: chainArgs(
			cobra.ExactArgs(1),
			func(_ *cobra.Command, args []string) error {
				opts.Path = args[0]
				return nil
			},
		),
		PreRunE: app.Setup(opts),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Run(cmd.Context(), async(func() error {
				return sign(opts.Path, opts.Signing)
			}))
		},
	}

	opts.AddFlags(cmd.Flags())
	commonConfiguration(cmd)

	return cmd
}

func sign(binPath string, opts options.Signing) error {
	err := validatePathIsDarwinBinary(binPath)
	if err != nil {
		return err
	}

	var cfg *quill.SigningConfig

	if opts.P12 != "" {
		cfg, err = quill.NewSigningConfigFromP12(binPath, opts.P12, opts.Password)
		if err != nil {
			return fmt.Errorf("unable to read p12: %w", err)
		}
	}

	cfg.WithIdentity(opts.Identity)
	cfg.WithTimestampServer(opts.TimestampServer)

	return quill.Sign(cfg)
}

func validatePathIsDarwinBinary(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}

	if _, err := macho.NewFile(fi); err != nil {
		return fmt.Errorf("given path=%q may not be a macho formatted binary: %w", path, err)
	}
	return err
}
