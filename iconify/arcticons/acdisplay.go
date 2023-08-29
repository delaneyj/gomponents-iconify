package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acdisplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.86 4.5a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27.2a2 2 0 0 0 2-2V30.83h-2.92v7.79H8.79V9.38h25.35v1.95h2.92V6.5a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.31 14.25a6.82 6.82 0 0 0-6.82 6.83v6.82h6.82a6.83 6.83 0 1 0 0-13.65Z"/>`),
		g.Group(children),
	)
}