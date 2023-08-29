package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24.157 44c11.046 0 20-8.954 20-20s-8.954-20-20-20s-20 8.954-20 20s8.954 20 20 20ZM17 14a1 1 0 0 0-1 1v.723l-2.468-2.107l-1.299 1.522L16 18.353V33a1 1 0 1 0 2 0V20.06l3 2.56V33a1 1 0 1 0 2 0v-8.673l3 2.56V33a1 1 0 1 0 2 0v-4.405l3 2.56V33a1 1 0 1 0 2 0v-.138l1.82 1.554l1.3-1.52L33 30.232V15a1 1 0 1 0-2 0v13.526l-3-2.56V15a1 1 0 1 0-2 0v9.258l-3-2.56V15a1 1 0 1 0-2 0v4.99l-3-2.56V15a1 1 0 0 0-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}