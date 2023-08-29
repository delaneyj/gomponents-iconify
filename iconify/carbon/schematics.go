package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schematics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 19.001A4.006 4.006 0 0 0 23 15H9a2.003 2.003 0 0 1-2-2V9.857A3.995 3.995 0 0 0 9.858 7h12.284a4 4 0 1 0 0-2H9.858A3.992 3.992 0 1 0 5 9.858v3.141A4.006 4.006 0 0 0 9.001 17H23a2.003 2.003 0 0 1 2 2.001V22h-3v3H9.858a4 4 0 1 0 0 2H22v3h8v-8h-3ZM26 4a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2ZM4 6a2 2 0 1 1 2 2a2.002 2.002 0 0 1-2-2Zm2 22a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm22-4v4h-4v-4Z"/>`),
		g.Group(children),
	)
}