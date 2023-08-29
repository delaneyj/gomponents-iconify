package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpOpenInBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4v16h6v-2H5V8h14v10h-4v2h6V4H3zm9 6l-4 4h3v6h2v-6h3l-4-4z"/>`),
		g.Group(children),
	)
}