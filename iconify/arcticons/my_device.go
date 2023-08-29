package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="25" x="4.5" y="11.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(90 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.596 21.568a6.5 6.5 0 1 0-9.192 0M17.5 29.472h13m-13 7h13"/>`),
		g.Group(children),
	)
}