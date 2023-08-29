package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="9.962" cy="24" r="4.956" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.918 24h13.536m-2.772 2.349V24m-6.606 2.5v13.542a3.458 3.458 0 0 0 3.459 3.458h17a3.458 3.458 0 0 0 3.459-3.458V7.958A3.458 3.458 0 0 0 39.535 4.5h-17a3.458 3.458 0 0 0-3.458 3.458V21.5m-.001-12.005h23.918m-23.918 29.01h23.918"/>`),
		g.Group(children),
	)
}