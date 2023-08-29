package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungBrowserLiteAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.448 25.21a14.718 14.718 0 0 1 26.233-10.6a14.718 14.718 0 0 1 .309 17.711m-3.2 3.446a14.716 14.716 0 0 1-23.002-6.163"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.802 23.671c20.812 19.897-54.214.628-27.462-7.359"/>`),
		g.Group(children),
	)
}