package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OsMaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.003 6.843c-.142 1.695-2.744 32.522-2.91 34.488a1.57 1.57 0 0 1-2.88.966c-.815-1.103-8.514-13.074-8.514-13.074L9.355 27.556a1.416 1.416 0 0 1-.61-2.671l28.68-19.401c1.525-1.032 2.719-.337 2.578 1.359Z"/><circle cx="30.722" cy="20.25" r="2.941" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}