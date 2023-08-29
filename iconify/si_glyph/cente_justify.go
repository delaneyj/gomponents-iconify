package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenteJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 1.938c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 1h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 4h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 7h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H1.98a.938.938 0 0 1 0-1.876h14.082c.518 0 .938.42.938.938zm-3 3c0 .518-.42.938-.938.938H4.98a.938.938 0 0 1 0-1.876h8.082c.518 0 .938.42.938.938z"/>`),
		g.Group(children),
	)
}