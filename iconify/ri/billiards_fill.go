package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BilliardsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm0 4a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm0 1.75a2.5 2.5 0 0 1 1.88 4.148c.565.456.92 1.117.92 1.852c0 1.38-1.254 2.5-2.8 2.5c-1.546 0-2.8-1.12-2.8-2.5c0-.735.355-1.396.92-1.853A2.5 2.5 0 0 1 12 7.75Zm0 5c-.754 0-1.3.488-1.3 1s.546 1 1.3 1s1.3-.488 1.3-1s-.546-1-1.3-1Zm0-3.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}