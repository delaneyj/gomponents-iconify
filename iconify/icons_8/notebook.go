package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 6v13.563l-2.28 2.312a2.452 2.452 0 0 0-.72 1.72A2.416 2.416 0 0 0 4.406 26h23.188A2.417 2.417 0 0 0 30 23.594c0-.64-.266-1.266-.72-1.72L27 19.564V6H5zm2 2h18v11H7V8zm-.563 13h19.125l2.313 2.28c.077.08.125.204.125.314c0 .24-.166.406-.406.406H4.406A.387.387 0 0 1 4 23.594a.47.47 0 0 1 .125-.313L6.438 21z"/>`),
		g.Group(children),
	)
}