package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drivedroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.21" cy="39.53" r="3.97" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.47" cy="17.93" r="3.97" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.21 4.5l-4.38 7.58h8.76L24.21 4.5zm8.18 9.87h7.11v7.11h-7.11zm-8.18 21.19V12.08"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.95 21.48v5.92H12.47v-5.5"/>`),
		g.Group(children),
	)
}