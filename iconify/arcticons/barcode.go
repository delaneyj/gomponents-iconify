package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="4.33" height="23.11" x="12.61" y="12.44" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".66"/><rect width="3.62" height="23.11" x="24.72" y="12.44" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".66"/><rect width="4.33" height="23.11" x="36.11" y="12.44" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".66"/><rect width="39" height="31" x="4.5" y="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.9 12.44v23.12m11.75-23.12v23.12M32.4 12.44v23.12"/>`),
		g.Group(children),
	)
}