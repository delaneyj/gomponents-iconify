package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CashLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm9.005-11.997h-18a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-14a1 1 0 0 0-1-1Zm-17 11.643V8.354a3.508 3.508 0 0 0 2.35-2.351h11.291a3.508 3.508 0 0 0 2.359 2.353v7.288a3.508 3.508 0 0 0-2.36 2.359H6.355a3.508 3.508 0 0 0-2.351-2.357Z"/>`),
		g.Group(children),
	)
}