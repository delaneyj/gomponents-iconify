package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.816 13.58c1.956 1.825 3.157 3.449 3.184 4.445m-3.428.593c-2.098-.634-4.944-2.03-7.919-3.976c-5.47-3.579-9.304-7.664-8.56-9.123c.32-.628 1.591-.6 3.294-.113"/><path d="M7.042 7.059a7 7 0 0 0 9.908 9.89m1.581-2.425A7 7 0 0 0 9.474 5.47M3 3l18 18"/></g>`),
		g.Group(children),
	)
}