package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kideapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m16.574 12.06l2.378-3.282l2.4 3.282v11.544l6.817-6.779h3.04v3.149l-8.46 8.413l-3.17.845l-3.018-3l.013-14.171Z"/><path d="m19.578 29.232l8.605 9.99h3.256v-3.135l-7.903-8.483"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}