package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tangerine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00D26A" d="M7 2h5c2.761 0 5 2.739 5 5.5h-5C9.239 7.5 7 4.761 7 2Z"/><path fill="#FF822D" d="M16.5 30C22.851 30 28 24.851 28 18.5S22.851 7 16.5 7S5 12.149 5 18.5S10.149 30 16.5 30Z"/></g>`),
		g.Group(children),
	)
}