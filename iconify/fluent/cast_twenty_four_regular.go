package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.25A2.25 2.25 0 0 1 4.25 4h15.5A2.25 2.25 0 0 1 22 6.25v11.5A2.25 2.25 0 0 1 19.75 20H4.25A2.25 2.25 0 0 1 2 17.75V6.25Zm2.25-.75a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h15.5a.75.75 0 0 0 .75-.75V6.25a.75.75 0 0 0-.75-.75H4.25Zm2.746 10.495a.998.998 0 1 1-1.996 0a.998.998 0 0 1 1.996 0Zm-2-3.244a.75.75 0 0 1 .75-.75a4.247 4.247 0 0 1 4.247 4.247a.75.75 0 0 1-1.5 0A2.747 2.747 0 0 0 5.746 13.5a.75.75 0 0 1-.75-.75Zm0-3.007a.75.75 0 0 1 .75-.75A7.254 7.254 0 0 1 13 16.248a.75.75 0 0 1-1.5 0a5.754 5.754 0 0 0-5.754-5.754a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}