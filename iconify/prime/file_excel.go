package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExcel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.53 9L13 3.47a.75.75 0 0 0-.53-.22H8A2.75 2.75 0 0 0 5.25 6v12A2.75 2.75 0 0 0 8 20.75h8A2.75 2.75 0 0 0 18.75 18V9.5a.75.75 0 0 0-.22-.5Zm-5.28-3.19l2.94 2.94h-2.94ZM16 19.25H8A1.25 1.25 0 0 1 6.75 18V6A1.25 1.25 0 0 1 8 4.75h3.75V9.5a.76.76 0 0 0 .75.75h4.75V18A1.25 1.25 0 0 1 16 19.25Z"/><path fill="currentColor" d="M14.47 11.91a.77.77 0 0 0-1.06.12L12 13.8L10.59 12a.75.75 0 1 0-1.18 1L11 15l-1.59 2a.75.75 0 0 0 1.18.94L12 16.2l1.41 1.8a.77.77 0 0 0 .59.28a.75.75 0 0 0 .59-1.28L13 15l1.63-2a.77.77 0 0 0-.16-1.09Z"/>`),
		g.Group(children),
	)
}