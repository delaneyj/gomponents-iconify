package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.877 7.502a6.625 6.625 0 1 1 13.25 0a6.625 6.625 0 0 1-13.25 0ZM1.85 7A5.676 5.676 0 0 1 7 1.849V4.5a.5.5 0 1 0 1 0V1.849A5.677 5.677 0 0 1 13.155 7H10.5a.5.5 0 0 0 0 1h2.655A5.676 5.676 0 0 1 8 13.155V10.5a.5.5 0 0 0-1 0v2.655A5.677 5.677 0 0 1 1.849 8H4.5a.5.5 0 0 0 0-1H1.849Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}