package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdvertisementLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.598 16l-3.2-8h-2l-2 5v.002L4.199 16h2.154l.4-1h3.29l.4 1h2.155Zm-5.046-3l.847-2.115L9.244 13H7.552ZM17 8h2v8h-3a3 3 0 1 1 0-6h1V8Zm-1 4a1 1 0 0 0 0 2h1v-2h-1Zm5-9H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1ZM4 19V5h16v14H4Z"/>`),
		g.Group(children),
	)
}