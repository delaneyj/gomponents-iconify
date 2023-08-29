package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bancolombia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.262 12.81c8.345-3.3 12.937-4.205 21.54-5.305M5.5 30.468c8.603-5.95 22.51-9.444 37-10.802M18.825 40.495c7.31-4.01 15.783-5.887 23.675-7.18"/>`),
		g.Group(children),
	)
}