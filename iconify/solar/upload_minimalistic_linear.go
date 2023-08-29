package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15c0 2.828 0 4.243.879 5.121C4.757 21 6.172 21 9 21h6c2.828 0 4.243 0 5.121-.879C21 19.243 21 17.828 21 15m-9 1V3m0 0l4 4.375M12 3L8 7.375"/>`),
		g.Group(children),
	)
}