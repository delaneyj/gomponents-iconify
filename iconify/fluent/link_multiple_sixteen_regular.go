package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMultipleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6.5A3.5 3.5 0 0 0 8.5 3h-4l-.192.005a3.5 3.5 0 0 0-1.305 6.66a4.566 4.566 0 0 1 .093-1.096A2.5 2.5 0 0 1 4.5 4h4l.164.006A2.5 2.5 0 0 1 8.5 9l-1.002.005l-.09.008a.5.5 0 0 0 .094.992l1-.005l.192-.005A3.5 3.5 0 0 0 12 6.5Zm2 3c0-.86-.435-1.62-1.096-2.07a4.517 4.517 0 0 0 .093-1.095a3.5 3.5 0 0 1-1.303 6.66l-.192.005l-1 .005a.51.51 0 0 1-.07-.005H7.5a3.5 3.5 0 0 1-.192-6.995L7.5 6h1a.5.5 0 0 1 .09.992L8.5 7h-1a2.5 2.5 0 0 0-.164 4.995L7.5 12H11v.002l.5-.002A2.5 2.5 0 0 0 14 9.5Z"/>`),
		g.Group(children),
	)
}