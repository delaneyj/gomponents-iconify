package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleBuzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 190q0 58-41 105L283 156l114-96q65 55 65 130zm-294 39L383 50Q317 5 232 5q-62 0-114 24T35 94zM64 316l91-76L25 109Q2 146 2 190q0 74 62 126zm345-8L270 167l-101 84l-2 2v1l-30 26h-1l-42 33l-16 14q6 4 10 6q-13 36-37 71q50 21 119-37q28 7 62 7q108 0 177-66z"/>`),
		g.Group(children),
	)
}