package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sahkoseuranta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.583 4.69l-11.42 25.357l9.725-1.47L14.245 43.5h7.468c0-.845 13.213-23.111 16.124-25.503c-4.49 0-12.263 1.518-12.263 1.518c2.85-5.325 6.73-10.02 10.118-15.015H21.685l-.101.19Z"/>`),
		g.Group(children),
	)
}