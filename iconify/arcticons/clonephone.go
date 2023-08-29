package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clonephone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="16.775" height="32.259" x="22.89" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.29"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.11 37.759v3.451a1.288 1.288 0 0 1-1.29 1.29H9.626a1.288 1.288 0 0 1-1.29-1.29V11.532a1.288 1.288 0 0 1 1.29-1.29H22.89M15.722 24.72H31.42m-11.933-2.807l-3.765 2.806m12.004 2.948l3.694-2.948"/>`),
		g.Group(children),
	)
}