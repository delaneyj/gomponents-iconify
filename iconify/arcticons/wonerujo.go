package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wonerujo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.69 28.846l.02-15.47m15.6 16.115c.081-5.299-.005-10.817-.005-16.116c-.209-3.57-.695-7.039-7.349-8.767M8.233 13.375v15.471c0 8.11 15.342 8.213 15.471 0c-.01 8.29 15.471 7.98 15.6.645m-31.086-.645v6.962"/><circle cx="39.032" cy="42.642" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}