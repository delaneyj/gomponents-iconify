package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M19 3H5c-1.104 0-1.99.896-1.99 2l-.008 3H3v11a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm0 8h-8v8H9v-8H5V9h4V5h2v4h8v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}