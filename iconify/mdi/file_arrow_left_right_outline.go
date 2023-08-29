package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileArrowLeftRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.09 20c.12.72.37 1.39.72 2H6c-1.11 0-2-.89-2-2V4a2 2 0 0 1 2-2h8l6 6v5.09c-.33-.05-.66-.09-1-.09c-.34 0-.67.04-1 .09V9h-5V4H6v16h7.09M23 17l-3-2.5V16h-4v2h4v1.5l3-2.5m-5 1.5L15 21l3 2.5V22h4v-2h-4v-1.5Z"/>`),
		g.Group(children),
	)
}