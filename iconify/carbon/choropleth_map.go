package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChoroplethMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.243 4.03l-8-2a1.006 1.006 0 0 0-.615.042l-9.7 3.88L3.243 4.03A1 1 0 0 0 2 5v22a1 1 0 0 0 .757.97l8 2A1.024 1.024 0 0 0 11 30a.995.995 0 0 0 .372-.072l9.7-3.88l7.686 1.922A1 1 0 0 0 30 27V5a1 1 0 0 0-.757-.97ZM28 11h-6V4.28l6 1.5Zm-18 8H4v-6h6Zm2-8V7.677l8-3.2V11Zm8 2v6h-8v-6Zm-8 8h8v3.323l-8 3.2Zm10-8h6v6h-6ZM10 7.78V11H4V6.28ZM4 21h6v6.72l-6-1.5Zm18 3.219V21h6v4.72Z"/>`),
		g.Group(children),
	)
}