package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLowSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 10v4a1.5 1.5 0 0 1-1.5 1.5h-.9v.31c0 1.14-.878 2.09-2.015 2.178l-1.268.099a68.39 68.39 0 0 1-10.634 0l-1.348-.105a2.392 2.392 0 0 1-2.173-1.99a24.149 24.149 0 0 1 0-7.985a2.392 2.392 0 0 1 2.173-1.989l1.348-.105a68.397 68.397 0 0 1 10.634 0l1.268.099A2.184 2.184 0 0 1 19.6 8.19v.31h.9A1.5 1.5 0 0 1 22 10ZM6.75 9.38a.75.75 0 0 0-.809-.748l-.234.019a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.236.02a.75.75 0 0 0 .811-.748V9.38Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}