package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpAutoClicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.864 19.278h10.858C36.979 9.9 32.609 4.11 25.616 4.52l.248 14.758Zm-3.405 0H11.602C11.345 9.9 15.715 4.11 22.708 4.52l-.249 14.758Zm-11.1 3.64C10.182 41.545 15.345 43.488 24 43.488c8.653 0 13.817-1.944 12.642-20.57H11.358Z"/>`),
		g.Group(children),
	)
}