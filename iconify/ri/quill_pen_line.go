package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuillPenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.94 14.033a29.79 29.79 0 0 0-.606 1.782c.96-.696 2.101-1.138 3.418-1.303c2.513-.314 4.746-1.974 5.876-4.059L14.172 9l1.413-1.415l1-1.002c.43-.429.915-1.224 1.428-2.367c-5.593.867-9.018 4.291-11.074 9.818ZM17 8.997l1 1c-1 3-4 6-8 6.5c-2.669.333-4.336 2.166-5.002 5.5H3c1-6 3-20 18-20c-1 2.997-1.998 4.996-2.997 5.997L17 8.997Z"/>`),
		g.Group(children),
	)
}