package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tophat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 1024H32q-13 0-22.5-9.5T0 992t9.5-22.5T32 960h160L128 85q-1-24 6-40t18.5-25T186 6.5T226.5 1T276 0h472q31 0 49.5 1T838 6.5T871.5 20t19 25t5.5 40l-64 875h160q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5zM243 832l13 128h512l14-128H243z"/>`),
		g.Group(children),
	)
}