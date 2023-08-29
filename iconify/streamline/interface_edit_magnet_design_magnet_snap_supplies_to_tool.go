package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditMagnetDesignMagnetSnapSuppliesToTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.81 12.07a4.86 4.86 0 0 1-6.88-6.88L6.62.5l2.19 2.19L4.51 7A1.77 1.77 0 0 0 7 9.49l4.3-4.3l2.2 2.19Zm.38-4.76l2.19 2.19M4.5 2.62l2.19 2.19"/>`),
		g.Group(children),
	)
}