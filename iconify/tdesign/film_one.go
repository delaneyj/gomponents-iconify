package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v2h2V4H4Zm4 0v7h8V4H8Zm10 0v2h2V4h-2Zm2 4h-2v3h2V8Zm0 5h-2v3h2v-3Zm0 5h-2v2h2v-2Zm-4 2v-7H8v7h8ZM6 20v-2H4v2h2Zm-2-4h2v-3H4v3Zm0-5h2V8H4v3Z"/>`),
		g.Group(children),
	)
}