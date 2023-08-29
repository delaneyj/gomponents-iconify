package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.728 28.702h4.702m-4.702-9.404h4.702M25.728 24h4.702m2.368 4.702H37.5m-4.702-9.404H37.5M32.798 24H37.5m-27-4.702h4.702M10.5 24h3.056M10.5 19.298v9.404m6.875 0v-9.404h3.079c1.74 0 3.151 1.414 3.151 3.158s-1.41 3.159-3.151 3.159h-3.08m3.08 0l3.078 3.085"/>`),
		g.Group(children),
	)
}