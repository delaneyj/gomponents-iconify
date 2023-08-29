package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myslt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.878 38.11c5.71-9.788 11.339-16.8 11.339-16.8c.955-1.416 4.548-.247 3.67 2.446a1633.653 1633.653 0 0 1-11.339 16.8c-1.481 2.262-5.235-.259-3.67-2.447ZM7.796 39.166c5.71-9.787 11.34-16.8 11.34-16.8c.955-1.415 4.547-.247 3.669 2.447a1633.653 1633.653 0 0 1-11.339 16.8c-1.48 2.261-5.234-.259-3.67-2.447ZM7.35 22.905c5.71-9.787 11.34-16.8 11.34-16.8c.955-1.415 4.548-.246 3.67 2.447a1633.653 1633.653 0 0 1-11.339 16.8c-1.481 2.262-5.235-.259-3.67-2.447Z"/><circle cx="27.137" cy="27.244" r="2.692" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}