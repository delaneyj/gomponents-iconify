package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Virb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M217 333H116L2 51h97l68 184l67-184h95zm245-233q0-20-13.5-34T415 52t-34 14t-14 34t14 33.5t34 13.5q19 0 33-13.5t14-33.5z"/>`),
		g.Group(children),
	)
}