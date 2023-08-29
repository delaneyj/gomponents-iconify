package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceMx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 5h-2l-2 3.897L16 5h-2l2.905 5L14 15h2l2-3.799L20 15h2l-2.902-5L22 5zM10 2L8.485 6.374L8 8l-.465-1.626L6 2H4v13h2V7.374l-.159-1.996l.58 1.996L8 12l1.579-4.626l.58-2l-.159 2V15h2V2h-2z"/><circle cx="9" cy="27" r="1" fill="currentColor"/><path fill="currentColor" d="M2 18h4v2H2zm6 0h4v2H8zm6 0h4v2h-4zm6 0h4v2h-4zm6 0h4v2h-4zm0 13H6a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM6 25v4h20v-4z"/>`),
		g.Group(children),
	)
}