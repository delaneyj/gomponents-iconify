package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h7.5v2H3V1Zm12 2.414V7h3.586L15 3.414ZM17.25 14.5a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5Zm-4.75 2.75a4.75 4.75 0 1 1 8.74 2.578l1.674 1.671l-1.413 1.415l-1.675-1.673A4.75 4.75 0 0 1 12.5 17.25Z"/>`),
		g.Group(children),
	)
}