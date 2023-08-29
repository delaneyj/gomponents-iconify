package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassCupOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.975 22q-.775 0-1.337-.5T5 20.225l-1.75-16q-.1-.875.488-1.55T5.224 2h13.55q.9 0 1.488.675t.487 1.55l-1.75 16q-.075.775-.637 1.275t-1.338.5H6.975Zm-.2-4L7 20h10l.225-2H6.775Zm-.2-2H17.45l1.3-12H5.25l1.325 12Zm.2 4h10.45h-10.45Z"/>`),
		g.Group(children),
	)
}