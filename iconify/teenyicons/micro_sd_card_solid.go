package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSdCardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h8A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V8.293l2-2V1.5ZM6 3v3h1V3H6Zm2 0v3h1V3H8Zm2 3V3h1v3h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}