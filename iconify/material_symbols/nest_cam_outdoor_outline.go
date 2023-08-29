package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamOutdoorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 19.275q-.55.65-1.363.738t-1.462-.463L7.35 16.025q-.65-.525-1.125-1.2t-.775-1.45q-.275.3-.65.463T4 14H2V6h2q.725 0 1.288.45t.687 1.175q.95-1.45 2.45-2.288T11.65 4.5q1.125 0 2.138.4t1.862 1.125L19.9 9.55q.65.525.713 1.338t-.463 1.462l-5.75 6.925Zm-7.4-8.3q-.025 1.025.413 1.938t1.212 1.562L12.85 18l5.75-6.925l-4.225-3.525q-.6-.5-1.325-.763t-1.5-.262Q9.675 6.5 8.35 7.8T7 10.975Zm5.775 1.275Z"/>`),
		g.Group(children),
	)
}