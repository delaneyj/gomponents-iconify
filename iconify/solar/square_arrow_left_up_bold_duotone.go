package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowLeftUpBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 3.464C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536Z" opacity=".5"/><path d="M8.421 9.172a.75.75 0 0 1 .75-.75h4.243a.75.75 0 1 1 0 1.5h-2.432l4.377 4.376a.75.75 0 0 1-1.061 1.061L9.92 10.982v2.433a.75.75 0 0 1-1.5 0V9.171Z"/></g>`),
		g.Group(children),
	)
}