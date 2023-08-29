package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M20 7.75a6 6 0 0 0-6 6v4.5a6 6 0 0 0 12 0v-4.5a6 6 0 0 0-6-6Zm-2.5 6a2.5 2.5 0 0 1 5 0v4.5a2.5 2.5 0 0 1-5 0v-4.5ZM12 9.5a1.75 1.75 0 0 0-3.017-1.207l-3.175 3.334a1.75 1.75 0 1 0 2.535 2.413l.157-.165V22.5a1.75 1.75 0 1 0 3.5 0v-13Z"/></g>`),
		g.Group(children),
	)
}