package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrScanTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3h6v6h-6V3ZM9 3v6H3V3h6Zm6 18v-6h6v6h-6Zm-6 0H3v-6h6v6ZM3 11h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}