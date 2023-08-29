package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EsetPaymentProtection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.27V10.4a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27.2a2 2 0 0 0 2 2h27.265M4.5 15.636h39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.488 41.6a11.308 11.308 0 0 0 7.913-9.394a38.052 38.052 0 0 0-.034-5.768a1.219 1.219 0 0 0-1.184-1.134a8.506 8.506 0 0 1-6.004-2.123a.984.984 0 0 0-1.381 0a8.506 8.506 0 0 1-6.004 2.123a1.219 1.219 0 0 0-1.184 1.134a38.052 38.052 0 0 0-.034 5.768a11.308 11.308 0 0 0 7.912 9.394Zm0 0V22.897"/>`),
		g.Group(children),
	)
}