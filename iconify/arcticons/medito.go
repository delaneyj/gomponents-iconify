package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medito(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.425 21.102s-1.293 2.196-2.696 2.196c-1.884 0-3.361-6.575-5.43-6.575c-1.478 0-5.356 2.18-5.8 6.021c3.03-.332 4.913 7.61 6.982 7.61c1.957 0 7.535-12.375 10.453-12.375c3.657 0 1.44 6.021 3.768 6.021c1.145 0 3.657-6.095 6.39-6.095c2.548 0 4.395 7.535 8.79 7.535a15.286 15.286 0 0 0 4.618-.53s-4.432 6.367-8.68 6.367S28.359 24 26.475 24a2.525 2.525 0 0 0-2.073 1.425"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.93 24.813c.185-.287.315-.74 1.008-.74c1.349 0-.17 7.204 2.743 7.204c2.549 0 7.61-8.902 7.61-8.902"/>`),
		g.Group(children),
	)
}