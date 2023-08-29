package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h5v2H4v3H2V2Zm9 0h5v5h-2V4h-3V2ZM9 9h13v13H9V9Zm2 2v9h9v-9h-9Zm-7-1v3h3v2H2v-5h2Z"/>`),
		g.Group(children),
	)
}