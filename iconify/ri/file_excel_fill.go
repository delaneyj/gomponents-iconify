package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExcelFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v14.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16Zm-2.8 10L16 8h-2.4L12 10.286L10.4 8H8l2.8 4L8 16h2.4l1.6-2.286L13.6 16H16l-2.8-4Z"/>`),
		g.Group(children),
	)
}