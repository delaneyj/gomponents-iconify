package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zmninja(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.67 23.63l41.11-8C33.4-6.39 2.86 1.37 2.67 23.63ZM9.7 25l7.62-1.55c.77 5.8-7.32 6.2-7.62 1.55Zm18.8-3.66l7.62-1.55c.77 5.83-7.32 6.21-7.62 1.55ZM4.41 32.68l40.92-8.07c.25 20.66-29.67 30.55-40.92 8.07Z"/>`),
		g.Group(children),
	)
}