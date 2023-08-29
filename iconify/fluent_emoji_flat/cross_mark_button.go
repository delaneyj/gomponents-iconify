package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossMarkButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00D26A" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M21.564 8.314a1.5 1.5 0 0 1 2.122 2.122L18.12 16l5.565 5.564a1.5 1.5 0 0 1-2.122 2.122L16 18.12l-5.564 5.565a1.5 1.5 0 0 1-2.122-2.122L13.88 16l-5.565-5.564a1.5 1.5 0 0 1 2.122-2.122L16 13.88l5.564-5.565Z"/></g>`),
		g.Group(children),
	)
}