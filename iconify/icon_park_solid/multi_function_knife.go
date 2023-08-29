package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiFunctionKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMultiFunctionKnife0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M30 10a6 6 0 0 0-12 0v28a6 6 0 0 0 12 0V10Zm0 10.314l7.243 7.242A4 4 0 0 0 42.9 21.9L30 9v11.314Z"/><path fill="#fff" stroke="#fff" d="m18.071 27.414l-7.243-7.242a4 4 0 1 0-5.656 5.656l12.9 12.9V27.414Z"/><path stroke="#000" stroke-linecap="round" d="M24 10v1m0 26v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMultiFunctionKnife0)"/>`),
		g.Group(children),
	)
}