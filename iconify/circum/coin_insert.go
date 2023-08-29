package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinInsert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.44 17.7h-3.67a7.484 7.484 0 0 0 1.78-4.86A7.55 7.55 0 1 0 6.23 17.7H2.56a.508.508 0 0 0-.5.5a.5.5 0 0 0 .5.5h18.88a.5.5 0 0 0 .5-.5a.508.508 0 0 0-.5-.5Zm-5.03 0H7.62a6.546 6.546 0 1 1 8.78-.01Z"/><path fill="currentColor" d="M14 13.965a1.616 1.616 0 0 1-1.5 1.61v.65a.485.485 0 0 1-.5.48a.491.491 0 0 1-.5-.48v-.64h-.81a.5.5 0 0 1-.5-.5a.508.508 0 0 1 .5-.5h1.69a.617.617 0 0 0 .62-.62a.623.623 0 0 0-.62-.62h-.75a1.618 1.618 0 0 1-.13-3.23v-.65a.491.491 0 0 1 .5-.48a.485.485 0 0 1 .5.48v.64h.81a.5.5 0 0 1 0 1h-1.68a.62.62 0 0 0 0 1.24h.75a1.626 1.626 0 0 1 1.62 1.62Z"/>`),
		g.Group(children),
	)
}