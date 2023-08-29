package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFiveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm2 5h-4a1 1 0 0 0-.993.883L9 8v4a1 1 0 0 0 .883.993L10 13h3v2h-2l-.007-.117A1 1 0 0 0 9 15a2 2 0 0 0 1.85 1.995L11 17h2a2 2 0 0 0 1.995-1.85L15 15v-2a2 2 0 0 0-1.85-1.995L13 11h-2V9h3a1 1 0 0 0 .993-.883L15 8a1 1 0 0 0-.883-.993L14 7z"/></g>`),
		g.Group(children),
	)
}