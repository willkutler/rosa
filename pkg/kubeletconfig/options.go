package kubeletconfig

import "github.com/spf13/cobra"

type KubeletConfigOptions struct {
	Name         string
	PodPidsLimit int
}

func NewKubeletConfigOptions() *KubeletConfigOptions {
	return &KubeletConfigOptions{}
}

func (k *KubeletConfigOptions) AddFlagsToCommand(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.SortFlags = false
	flags.IntVar(
		&k.PodPidsLimit,
		PodPidsLimitOption,
		PodPidsLimitOptionDefaultValue,
		PodPidsLimitOptionUsage)
	flags.StringVar(
		&k.Name,
		NameOption,
		NameOptionDefaultValue,
		NameOptionUsage)
}