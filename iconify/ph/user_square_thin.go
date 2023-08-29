package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a12 12 0 0 0-12 12v160a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12ZM63 212a68 68 0 0 1 130 0Zm149-4a4 4 0 0 1-4 4h-6.66a76 76 0 0 0-52.75-53.13a44 44 0 1 0-41.18 0A76 76 0 0 0 54.66 212H48a4 4 0 0 1-4-4V48a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Zm-84-52a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}