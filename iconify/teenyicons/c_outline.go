package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func COutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m10 5.5l-.068-.068a3.182 3.182 0 0 0-2.25-.932H7.5a3 3 0 0 0 0 6h.182c.844 0 1.653-.335 2.25-.932L10 9.5m-8.5 1v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`),
		g.Group(children),
	)
}