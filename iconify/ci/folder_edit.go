package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9.25V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H3m0 0v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874C4.52 20 5.08 20 6.2 20H7M3 6v-.4c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C3.76 4 4.04 4 4.6 4h4.737c.245 0 .367 0 .482.028a1 1 0 0 1 .29.12c.1.061.187.148.36.32L12 6m4 8l2 2m-7 5v-2.5l7.5-7.5l2.5 2.5l-7.5 7.5H11Z"/>`),
		g.Group(children),
	)
}