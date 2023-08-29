package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 6v13.563l-2.281 2.312A2.451 2.451 0 0 0 2 23.594A2.417 2.417 0 0 0 4.406 26h23.188A2.417 2.417 0 0 0 30 23.594c0-.64-.266-1.266-.719-1.719L27 19.562V6zm2 2h18v11H7zm-.563 13h19.125l2.313 2.281a.465.465 0 0 1 .125.313a.385.385 0 0 1-.406.406H4.406A.385.385 0 0 1 4 23.594c0-.11.047-.235.125-.313z"/>`),
		g.Group(children),
	)
}