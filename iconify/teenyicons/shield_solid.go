package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.748.066a.5.5 0 0 0-.496 0l-7 4A.5.5 0 0 0 0 4.5v.72a10.15 10.15 0 0 0 7.363 9.76a.5.5 0 0 0 .274 0A10.152 10.152 0 0 0 15 5.22V4.5a.5.5 0 0 0-.252-.434l-7-4Z"/>`),
		g.Group(children),
	)
}