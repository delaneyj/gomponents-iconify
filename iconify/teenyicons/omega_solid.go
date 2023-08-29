package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OmegaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 1 1 9 6.002V14h5v1H9v-3h.026a.499.499 0 0 1 .307-.313a5.5 5.5 0 1 0-3.667 0a.496.496 0 0 1 .308.313H6v3H0v-1h5v-1.498A6.502 6.502 0 0 1 1 6.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}