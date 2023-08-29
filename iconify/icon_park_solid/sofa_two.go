package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SofaTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSofaTwo0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 38V18h-8v13H12V18H4v20h40Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 6H12v25h24V6Z"/><path fill="#fff" fill-rule="evenodd" d="M10 44a4 4 0 0 0 4-4c-1.097.004-7.3 0-8 0a4 4 0 0 0 4 4Zm28 0a4 4 0 0 0 4-4c-1.905-.007-7.137 0-8 0a4 4 0 0 0 4 4Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSofaTwo0)"/>`),
		g.Group(children),
	)
}