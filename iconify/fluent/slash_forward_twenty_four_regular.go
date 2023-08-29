package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.499 2.042a.75.75 0 0 1 .459.957l-6.5 18.5A.75.75 0 0 1 8.043 21l6.5-18.5a.75.75 0 0 1 .956-.459Z"/>`),
		g.Group(children),
	)
}