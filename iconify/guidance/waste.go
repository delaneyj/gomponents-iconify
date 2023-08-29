package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.5 4C12.5 2 14 .5 16 .5m-4 6.08l-.31-.196c-1.441-.912-3.266-1.258-4.717-.36C6.243 6.473 5.079 7.375 4.5 8v.293c2.08 1.073 4 3.226 4 5.707s-1.92 4.634-4 5.706V20c.58.624 1.743 1.525 2.473 1.977c1.45.897 3.276.551 4.717-.36l.31-.197l.31.196c1.441.912 3.266 1.258 4.717.36c.73-.45 1.894-1.352 2.473-1.976v-.294c-2.08-1.072-4-3.225-4-5.706c0-2.481 1.92-4.634 4-5.707V8c-.58-.624-1.743-1.525-2.473-1.977c-1.45-.897-3.276-.551-4.717.36L12 6.58Z"/>`),
		g.Group(children),
	)
}