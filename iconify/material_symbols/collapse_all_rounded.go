package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseAllRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17.4l-3.9 3.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l3.875-3.875q.575-.575 1.425-.575t1.425.575L17.3 19.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L12 17.4Zm0-10.8l3.9-3.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-3.875 3.875Q12.85 8.55 12 8.55t-1.425-.575L6.7 4.1q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L12 6.6Z"/>`),
		g.Group(children),
	)
}