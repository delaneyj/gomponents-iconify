package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14V5c-1.5.167-5 0-7-2c-.571.571-1.265.993-2 1.3M5 5v9c0 4 7 7 7 7s3.204-1.373 5.277-3.5M3 3l18 18"/>`),
		g.Group(children),
	)
}