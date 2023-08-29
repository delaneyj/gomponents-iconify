package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1.938c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 1h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 4h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 7h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 10h14.082c.518 0 .938.42.938.938zm0 3c0 .518-.42.938-.938.938H5.98a.938.938 0 0 1 0-1.876h9.082c.518 0 .938.42.938.938z"/>`),
		g.Group(children),
	)
}