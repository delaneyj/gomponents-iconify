package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.372 20.222v-5.358h1.79V22H4v-7.136h1.79v5.358h12.582Z"/><path fill="currentColor" d="m7.768 14.356l8.79 1.824l.372-1.755L8.14 12.6l-.372 1.756Zm1.162-4.157l8.14 3.764l.744-1.617l-8.14-3.787l-.744 1.64Zm2.256-3.973l6.907 5.705l1.14-1.363l-6.907-5.704l-1.14 1.362ZM15.651 2L14.21 3.062l5.35 7.16L21 9.159L15.651 2Zm-8.07 16.42h8.977v-1.778H7.581v1.778Z"/>`),
		g.Group(children),
	)
}