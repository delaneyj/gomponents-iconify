package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookArrowRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.09 20c.12.72.37 1.39.72 2H6c-1.11 0-2-.89-2-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v9.09c-.33-.05-.66-.09-1-.09c-.34 0-.67.04-1 .09V4h-5v8l-2.5-2.25L8 12V4H6v16h7.09M23 19l-3-3v2h-4v2h4v2l3-3Z"/>`),
		g.Group(children),
	)
}