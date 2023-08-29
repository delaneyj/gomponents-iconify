package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 2.5A4.25 4.25 0 0 1 11 6.75V11H6.75a4.25 4.25 0 0 1 0-8.5Zm0 10.5H11v4.25A4.25 4.25 0 1 1 6.75 13Zm10.5-10.5a4.25 4.25 0 0 1 0 8.5H13V6.75a4.25 4.25 0 0 1 4.25-4.25ZM13 13h4.25A4.25 4.25 0 1 1 13 17.25V13Z"/>`),
		g.Group(children),
	)
}