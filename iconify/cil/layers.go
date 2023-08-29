package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m487.938 162.108l-224-128a16 16 0 0 0-15.876 0l-224 128a16 16 0 0 0 .382 28l224 120a16 16 0 0 0 15.112 0l224-120a16 16 0 0 0 .382-28ZM256 277.849L65.039 175.548L256 66.428l190.961 109.12Z"/><path fill="currentColor" d="M263.711 394.02L480 275.061v-36.522L256 361.74L32 238.539v36.522L248.289 394.02a16.005 16.005 0 0 0 15.422 0Z"/><path fill="currentColor" d="m32 362.667l216.471 115.451a16 16 0 0 0 15.058 0L480 362.667V326.4L256 445.867L32 326.4Z"/>`),
		g.Group(children),
	)
}