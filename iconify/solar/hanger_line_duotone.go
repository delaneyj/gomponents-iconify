package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangerLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9.536 6.91c0-1.055.95-1.91 2.124-1.91c1.173 0 2.124.855 2.124 1.91c0 .495-.18.947-.492 1.287c-.597.65-1.49 1.305-1.49 2.149v.284m0 0a3.656 3.656 0 0 1 2.082.61l7.433 5.01c1.306.881.613 2.75-1.02 2.75H3.703c-1.616 0-2.32-1.835-1.048-2.73l7.096-4.991a3.641 3.641 0 0 1 2.05-.65Z"/>`),
		g.Group(children),
	)
}