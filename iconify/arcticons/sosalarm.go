package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sosalarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.636 32.205h8.097m-8.097-13.974l4.044-2.229m0 0v16.203m6.084-10.937A5.407 5.407 0 0 1 34.03 15.8a5.443 5.443 0 0 1 3.849 9.318c-2.228 1.824-9.113 7.09-9.113 7.09h10.739m-31-.003h8.1m-8.1-13.974l4.047-2.229m0 0v16.203"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}