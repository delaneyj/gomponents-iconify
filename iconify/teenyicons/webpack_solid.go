package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebpackSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m8 4.217l2.25 1.35l3.268-1.635L8 .712v3.505ZM7 .713L1.482 3.932L4.75 5.566L7 4.216V.714ZM1 4.809v5.422l3.5-1.556V6.56L1 4.809Zm.523 6.283L7 14.287v-3.504l-2.034-1.22l-3.443 1.53ZM8 14.287l5.477-3.195l-3.443-1.53L8 10.783v3.504Zm6-4.057V4.81l-3.5 1.75v2.116l3.5 1.556Zm-6-.613l1.5-.9V7.059l-1.5.75v1.808ZM7 7.809v1.808l-1.5-.9V7.059l1.5.75ZM5.811 6.096l1.689.845l1.689-.845L7.5 5.083L5.811 6.096Z"/>`),
		g.Group(children),
	)
}