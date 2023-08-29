package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdvertisementFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.552 13l.847-2.115L9.244 13H7.552ZM16 12h1v2h-1a1 1 0 1 1 0-2Zm5-9H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Zm-8.402 13h-2.155l-.4-1h-3.29l-.4 1H4.199l1.199-2.998l.001-.002l2-5h2l3.199 8ZM17 8h2v8h-3a3 3 0 1 1 0-6h1V8Z"/>`),
		g.Group(children),
	)
}