package example

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerCreate,
		ReadContext:   resourceServerRead,
		UpdateContext: resourceServerUpdate,
		DeleteContext: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_ssh_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type Mounts struct {
	id   uuid.UUID
	name string
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id1, _ := uuid.NewUUID()
	d.SetId(id1.String())

	//newComponent := environment.InstalledComponentVersion{
	//	EnvironmentRef: id1,
	//	Name:           "name",
	//	Type:           "1",
	//	Version:        "1",
	//	Image:          "1",
	//	WorkDirectory:  "1",
	//	Mounts:         []string{"mount"},
	//	Shared:         "1",
	//}
	//
	//config, err := configuration.GetConfig()
	//e, err := environment.Get(config.CurrentEnvironment)
	//
	//err = e.Install(newComponent)
	//
	//if err != nil {
	//
	//}newComponent := environment.InstalledComponentVersion{
	//	EnvironmentRef: id1,
	//	Name:           "name",
	//	Type:           "1",
	//	Version:        "1",
	//	Image:          "1",
	//	WorkDirectory:  "1",
	//	Mounts:         []string{"mount"},
	//	Shared:         "1",
	//}
	//
	//config, err := configuration.GetConfig()
	//e, err := environment.Get(config.CurrentEnvironment)
	//
	//err = e.Install(newComponent)
	//
	//if err != nil {
	//
	//}

	//cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	//if err != nil {
	//	panic(err)
	//}
	//
	////reader, err := cli.ImagePull(ctx, "docker.io/epiphanyplatform/epicli:0.9.0", types.ImagePullOptions{})
	////if err != nil {
	////	panic(err)
	////}
	////io.Copy(os.Stdout, reader)
	//
	//resp, err := cli.ContainerCreate(ctx, &container.Config{
	//	Image: "docker.io/epiphanyplatform/epicli:0.9.0",
	//	Cmd:   []string{"--help"},
	//	Tty:   false,
	//}, nil, nil, nil, "")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	//	panic(err)
	//}
	//
	//statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	//select {
	//case err := <-errCh:
	//	if err != nil {
	//		panic(err)
	//	}
	//case <-statusCh:
	//}
	//
	//out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Print(out)

	var diagnostics diag.Diagnostics
	resourceServerRead(ctx, d, m)
	return diagnostics
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	return diagnostics
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resourceServerRead(ctx, d, m)

	var diagnostics diag.Diagnostics
	return diagnostics
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	return diagnostics
}
