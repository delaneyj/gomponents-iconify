package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wikis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm12 13h-6a24.26 24.26 0 0 0-2.79-10.55A12 12 0 0 1 28 15ZM16 28a5 5 0 0 1-.67 0A21.85 21.85 0 0 1 12 17h8a21.85 21.85 0 0 1-3.3 11a5 5 0 0 1-.7 0Zm-4-13a21.85 21.85 0 0 1 3.3-11a6 6 0 0 1 1.34 0A21.85 21.85 0 0 1 20 15Zm.76-10.55A24.26 24.26 0 0 0 10 15H4a12 12 0 0 1 8.79-10.55ZM4.05 17h6a24.26 24.26 0 0 0 2.75 10.55A12 12 0 0 1 4.05 17Zm15.16 10.55A24.26 24.26 0 0 0 22 17h6a12 12 0 0 1-8.79 10.55Z"/>`),
		g.Group(children),
	)
}