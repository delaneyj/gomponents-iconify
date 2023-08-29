package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M515 1024q-127 0-239.5-59.5T90 802T0 576q57 86 153 139t202 53q113 0 208.5-55.5T714 561t55-209q0-106-53.5-200.5T576 0q123 17 225.5 89T964 273t60 239q0 104-40.5 199T875 874.5t-162.5 109T515 1024z"/>`),
		g.Group(children),
	)
}