package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.437 18.859H4.563a2.5 2.5 0 0 1-2.5-2.5V7.641a2.5 2.5 0 0 1 2.5-2.5h14.874a2.5 2.5 0 0 1 2.5 2.5v8.718a2.5 2.5 0 0 1-2.5 2.5ZM4.563 6.141a1.5 1.5 0 0 0-1.5 1.5v8.718a1.5 1.5 0 0 0 1.5 1.5h14.874a1.5 1.5 0 0 0 1.5-1.5V7.641a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M8.063 14.247h-3a.5.5 0 1 1 0-1h3a.5.5 0 1 1 0 1Zm10.871.002h-6.5a.5.5 0 0 1 0-1h6.5a.5.5 0 0 1 0 1Z"/><rect width="2" height="4" x="16.434" y="7.14" fill="currentColor" rx=".5" transform="rotate(-90 17.434 9.14)"/>`),
		g.Group(children),
	)
}