package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.657c-2.498-4.025-5.529-7.102-9.021-7.102S4.5 16.182 4.5 27.007c0 7.98 2.151 9.438 4.441 9.438c5.899 0 12.168-12.19 15.059-17.788Zm0 0c2.498-4.025 5.529-7.102 9.021-7.102S43.5 16.182 43.5 27.007c0 7.98-2.151 9.438-4.441 9.438c-5.899 0-12.168-12.19-15.059-17.788Z"/>`),
		g.Group(children),
	)
}