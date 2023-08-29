package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorCareAnywhere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.268 45.406c-2.001 0-4.01-.288-5.97-.857a4.5 4.5 0 0 1-3.066-5.576a4.498 4.498 0 0 1 5.577-3.067c1.143.332 2.307.5 3.459.5c6.84 0 12.406-5.565 12.406-12.406s-5.566-12.406-12.406-12.406S11.86 17.159 11.86 24c0 1.385.244 2.773.723 4.129A4.5 4.5 0 0 1 4.1 31.13A21.343 21.343 0 0 1 2.86 24c0-11.804 9.603-21.406 21.407-21.406S45.674 12.196 45.674 24S36.07 45.406 24.268 45.406Z"/><circle cx="24.268" cy="24" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}