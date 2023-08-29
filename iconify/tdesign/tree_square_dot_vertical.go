package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeSquareDotVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4h3v3h-3V4Zm-2 2.5V9h7V2h-7v2.5h-3.667V11H9V8.5H2v7h7V13h2.333v6.5H15V22h7v-7h-7v2.5h-1.667v-11H15ZM17 20v-3h3v3h-3ZM7 10.5v3H4v-3h3Z"/>`),
		g.Group(children),
	)
}