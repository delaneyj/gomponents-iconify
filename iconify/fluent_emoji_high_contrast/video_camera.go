package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M11 19a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-2Z"/><path d="M9 17.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5v.5h1a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-1v.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 9 22.5v-5Zm11 0a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-5Z"/><path d="M3 8a1 1 0 0 1 1 1h1.5A1.5 1.5 0 0 1 7 10.5V10h11a4.002 4.002 0 0 1 4 4h7a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-3v6a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V13.5A1.5 1.5 0 0 1 5.5 15H4a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1Zm25 13v-6h-2v6h2Zm-3-6H8v10h17V15ZM13 27.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 0-1h-1a.5.5 0 0 0-.5.5Zm5.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.33-15A3.001 3.001 0 0 0 18 11H8v2h12.83Z"/></g>`),
		g.Group(children),
	)
}