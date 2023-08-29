package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fdroidbuildstatus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.6 33.52l2.9 2.37l3.5-1.33l.6-3.7l-2.9-2.37m6.7 9.67l-4.4-3.6"/><circle cx="16.99" cy="15.18" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.01" cy="15.18" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="31.32" height="12.99" x="8.35" y="8.68" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.83L9.16 9.5"/><rect width="31.32" height="17.31" x="8.35" y="24.67" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.5 5.85l-3.66 3.67"/>`),
		g.Group(children),
	)
}