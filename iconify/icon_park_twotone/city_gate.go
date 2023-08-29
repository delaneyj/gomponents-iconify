package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCityGate0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M15 11h17s5.048 2.966 7 4c1.091.578 5 1 5 1s-1.816.649-3 1c-1.544.458-4 1-4 1H11s-2.456-.542-4-1c-1.184-.351-3-1-3-1s3.909-.422 5-1c1.952-1.034 6-4 6-4Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 11l7-7l6 7H17Z"/><path d="M35 18v6m-23-6v6"/><path stroke-linejoin="round" d="m4 44l2-20h36l2 20H4Z"/><path d="M20 38a4 4 0 0 1 8 0v6h-8v-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCityGate0)"/>`),
		g.Group(children),
	)
}