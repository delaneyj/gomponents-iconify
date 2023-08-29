package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.75 4a.75.75 0 0 0-1.5 0v1.816a3.375 3.375 0 0 0-2.872 2.899l-.087.653c-.015.11-.029.221-.042.332a.493.493 0 0 0 .492.55H20.26a.493.493 0 0 0 .492-.55c-.014-.11-.027-.222-.042-.332l-.087-.653a3.375 3.375 0 0 0-2.872-2.899V4a.75.75 0 0 0-1.5 0v1.668a47.911 47.911 0 0 0-8.5 0V4Zm13.195 8.226a.494.494 0 0 0-.496-.476H3.551a.494.494 0 0 0-.496.476a28.92 28.92 0 0 0 .33 5.41a3.01 3.01 0 0 0 2.678 2.532l1.193.118c3.155.31 6.333.31 9.488 0l1.193-.118a3.01 3.01 0 0 0 2.678-2.532a28.99 28.99 0 0 0 .33-5.41Z"/>`),
		g.Group(children),
	)
}