package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batterycalibration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.98 12.31h1.93v1.93h-1.93zm8.1 0h1.93v1.93h-1.93z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.95 12.31V9.25h-6.79V24H2.5m25.55-11.69V9.25h6.79V24H45.5"/><rect width="13.84" height="24.96" x="17.08" y="14.24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.37 18.22h0l-4.98 10.35l7.82-2.2l-4.99 8.85l4.99-8.85"/>`),
		g.Group(children),
	)
}