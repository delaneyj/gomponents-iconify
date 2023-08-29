package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 2.5A4.25 4.25 0 0 1 11 6.75V11H6.75a4.25 4.25 0 0 1 0-8.5ZM9 9V6.75A2.25 2.25 0 1 0 6.75 9H9Zm-2.25 4H11v4.25A4.25 4.25 0 1 1 6.75 13Zm0 2A2.25 2.25 0 1 0 9 17.25V15H6.75Zm10.5-12.5a4.25 4.25 0 0 1 0 8.5H13V6.75a4.25 4.25 0 0 1 4.25-4.25Zm0 6.5A2.25 2.25 0 1 0 15 6.75V9h2.25ZM13 13h4.25A4.25 4.25 0 1 1 13 17.25V13Zm2 2v2.25A2.25 2.25 0 1 0 17.25 15H15Z"/>`),
		g.Group(children),
	)
}