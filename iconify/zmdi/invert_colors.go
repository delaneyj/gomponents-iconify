package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertColors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M291 121q50 50 50 121t-50 120.5T170.5 412T50 362.5T0 242t50-121L171 0zM171 370V61l-91 90q-37 38-37 91t37 90q37 38 91 38z"/>`),
		g.Group(children),
	)
}