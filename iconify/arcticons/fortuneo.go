package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fortuneo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.56 8.58l-2.45-.18l-2.76-3.9l-4.63 3l.49 3.19l-1.74 1.5ZM18.19 21.11L34.48 43.5l-5.41-27.93H6Zm0 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.93 43.5l10.21-14.2l-5.95-8.19ZM42 15.57l-11.89 5.37l-1-5.37Zm0 0"/>`),
		g.Group(children),
	)
}