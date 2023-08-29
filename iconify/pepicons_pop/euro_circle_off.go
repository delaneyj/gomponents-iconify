package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.489 7C12.43 7 10 9.213 10 12.5c0 3.386 2.527 6 5.489 6c.743 0 1.451-.161 2.098-.454a1 1 0 1 1 .826 1.821a7.061 7.061 0 0 1-2.924.633C11.283 20.5 8 16.846 8 12.5C8 8.055 11.38 5 15.489 5c1.237 0 2.428.393 3.574 1.174a1 1 0 1 1-1.126 1.652C17.079 7.242 16.274 7 15.489 7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 11a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Zm0 3.5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}