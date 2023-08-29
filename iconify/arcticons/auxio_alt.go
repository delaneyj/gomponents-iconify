package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuxioAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.348 4.5h-4.64v22.613l.004.005a8.693 8.693 0 1 0 4.636 7.687v-21.61h8.695V4.5Z"/>`),
		g.Group(children),
	)
}