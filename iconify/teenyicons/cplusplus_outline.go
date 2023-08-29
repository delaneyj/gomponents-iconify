package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CplusplusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 9.5c-.64.64-1.509 1-2.414 1H6.5a3 3 0 0 1 0-6h.586c.905 0 1.774.36 2.414 1m-2 .5v3M6 7.5h6M10.5 6v3m-9 1.5v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`),
		g.Group(children),
	)
}