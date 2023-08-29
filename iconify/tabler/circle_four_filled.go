package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm2 5a1 1 0 0 0-.993.883L13 8v3h-2V8l-.007-.117a1 1 0 0 0-1.986 0L9 8v3l.005.15a2 2 0 0 0 1.838 1.844L11 13h2v3l.007.117a1 1 0 0 0 1.986 0L15 16V8l-.007-.117A1 1 0 0 0 14 7z"/></g>`),
		g.Group(children),
	)
}