package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindmillOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.061 11.06C16.241 10.236 17 8.95 17 7.5C17 5.01 14.76 3 12 3v5m0 4c0 2.76 2.01 5 4.5 5c.166 0 .33-.01.49-.03m2.624-1.36C20.47 14.7 21 13.42 21 12h-5m-4 0c-2.76 0-5 2.01-5 4.5S9.24 21 12 21v-9zM6.981 7.033C4.737 7.318 3 9.435 3 12h9M3 3l18 18"/>`),
		g.Group(children),
	)
}