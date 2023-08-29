package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M0 12v3h3l8-8l-3-3l-8 8zm3 2H1v-2h1v1h1v1zm10.3-9.3L12 6L9 3l1.3-1.3a.996.996 0 0 1 1.41 0l1.59 1.59c.39.39.39 1.02 0 1.41z" fill="currentColor"/>`),
		g.Group(children),
	)
}