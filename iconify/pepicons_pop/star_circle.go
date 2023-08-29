package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.705 20.886L13 18.97l4.295 1.915a1 1 0 0 0 1.402-1.019l-.494-4.677l3.148-3.493a1 1 0 0 0-.535-1.647l-4.6-.976L13.865 5a1 1 0 0 0-1.732 0l-2.35 4.074l-4.6.976a1 1 0 0 0-.535 1.647l3.148 3.494l-.494 4.676a1 1 0 0 0 1.402 1.018Zm3.888-3.924l-3.119 1.39l.359-3.395a1 1 0 0 0-.252-.774l-2.286-2.537l3.34-.708a1 1 0 0 0 .66-.478L13 7.502l1.706 2.958a1 1 0 0 0 .659.478l3.34.708l-2.286 2.537a1 1 0 0 0-.252.774l.359 3.396l-3.119-1.39a1 1 0 0 0-.814 0Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}