package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.934 11.751l-4.75-8.255a.501.501 0 0 0-.434-.25h-9.5a.501.501 0 0 0-.434.25l-4.75 8.255a.498.498 0 0 0 0 .498l4.75 8.255c.09.155.255.25.434.25h9.5a.501.501 0 0 0 .434-.25l4.75-8.255a.498.498 0 0 0 0-.498zm-5.473 8.004H7.539L3.077 12L7.54 4.245h8.922L20.923 12l-4.462 7.755z"/>`),
		g.Group(children),
	)
}