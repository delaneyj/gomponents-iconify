package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSettingsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v14.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16ZM8.595 12.812l-.992.572l1 1.732l.992-.573c.393.372.873.654 1.405.812V16.5h2v-1.145a3.495 3.495 0 0 0 1.405-.812l.992.573l1-1.732l-.991-.573a3.51 3.51 0 0 0 0-1.623l.991-.572l-1-1.732l-.992.573A3.496 3.496 0 0 0 13 8.645V7.5h-2v1.145a3.496 3.496 0 0 0-1.405.812l-.992-.573l-1 1.732l.992.573a3.51 3.51 0 0 0 0 1.623ZM12 13.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}