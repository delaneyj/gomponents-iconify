package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 4a7 7 0 1 1-7 7a7 7 0 0 1 7-7m0-2a9 9 0 1 0 9 9a9 9 0 0 0-9-9Z"/><path fill="currentColor" d="M28 11a13.956 13.956 0 0 0-4.105-9.895L22.48 2.52a11.994 11.994 0 0 1-16.924 17l-.038-.038l-1.414 1.414A13.956 13.956 0 0 0 14 25v3h-4v2h10v-2h-4v-3.16A14.01 14.01 0 0 0 28 11Z"/>`),
		g.Group(children),
	)
}