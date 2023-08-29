package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hedera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0a12 12 0 1 0 0 24a12 12 0 0 0 0-24Zm4.957 17.396h-1.581V14.01H8.622v3.378H7.05V6.604h1.58v3.384h6.753V6.604h1.582zm-1.581-6.259H8.622v1.724h6.754Z"/>`),
		g.Group(children),
	)
}