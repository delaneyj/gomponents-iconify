package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidNineteenVirusPandemicTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.885 6.648a5.352 5.352 0 0 1 0 10.7m-.892-14.713h1.784m-.892 0v4.013m5.991-1.901l1.262 1.261m-.631-.63l-2.838 2.838m5.581 2.892v1.784m0-.892h-4.014m1.902 5.992l-1.262 1.261m.631-.631l-2.838-2.838m-2.892 5.581h-1.784m.892 0v-4.013M10.25 21.5a9.5 9.5 0 0 1 0-19v19Z"/><path d="M10.25 2.5A10.924 10.924 0 0 0 5.5 12a10.924 10.924 0 0 0 4.75 9.5m0-12.571H1.665M10 15.071H1.415"/></g>`),
		g.Group(children),
	)
}