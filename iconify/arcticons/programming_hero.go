package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingHero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.912 7.493L8.24 22.165m11.207 5.771l-7.293 7.292m7.3-18.902l-8.562 8.562m8.333-2.286L6.73 35.097m10.683-.514c2.435-2.435 5.48-8.44 12.99-8.44S41.27 21.367 41.27 15.366a10.871 10.871 0 0 0-21.742-.022l-.008 25.413a2.7 2.7 0 0 0 1.332 2.387a2.66 2.66 0 0 0 3.99-2.305V27.937"/>`),
		g.Group(children),
	)
}