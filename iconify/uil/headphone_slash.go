package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.07 5.936a9.945 9.945 0 0 0-9.39-2.671a1 1 0 0 0 .44 1.95A8.02 8.02 0 0 1 20 13h-.34a1 1 0 0 0 0 2H20v.34a1 1 0 1 0 2 0V13a9.888 9.888 0 0 0-2.93-7.064ZM3.706 2.293a1 1 0 0 0-1.414 1.414l2.435 2.435A9.962 9.962 0 0 0 2 13v7a1 1 0 0 0 1 1h3a3.003 3.003 0 0 0 3-3v-2a3.003 3.003 0 0 0-3-3H4a7.963 7.963 0 0 1 2.145-5.441L15 16.414V18a3.003 3.003 0 0 0 3 3h1.586l.707.707a1 1 0 0 0 1.414-1.414ZM6 15a1.001 1.001 0 0 1 1 1v2a1.001 1.001 0 0 1-1 1H4v-4Z"/>`),
		g.Group(children),
	)
}