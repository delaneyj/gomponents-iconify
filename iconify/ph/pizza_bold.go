package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PizzaBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M243.43 62.05a19.93 19.93 0 0 0-9.06-12.38a205.51 205.51 0 0 0-212.74 0a20 20 0 0 0-6.7 27.48l96 157.26a20 20 0 0 0 34.2 0l96-157.26a19.82 19.82 0 0 0 2.3-15.1ZM87 149.18l-19.87-32.56A19.82 19.82 0 0 1 72 116a20 20 0 0 1 15 33.18Zm64 29.45A20 20 0 0 1 168 148c.56 0 1.12 0 1.67.07Zm31.78-52.08a44 44 0 0 0-44.9 73.57L128 216.36L99.77 170.1a44 44 0 0 0-20.28-77.45a140.2 140.2 0 0 1 118 9.81ZM210 82a164.15 164.15 0 0 0-164 0l-8.45-13.86a181.52 181.52 0 0 1 180.9 0Z"/>`),
		g.Group(children),
	)
}