package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdealoShopping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.048 29.602H22.615c-.645 0-1.142-.57-1.054-1.21l2.481-18.048a.613.613 0 0 0-.607-.696h-3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.752 12.453h17.955c.47 0 .838.408.789.876l-1.002 9.55a2.024 2.024 0 0 1-2.012 1.812H22.07m-4.491-8.185H12.07m4.546 6.438H5.5"/><ellipse cx="21.66" cy="35.526" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.086" ry="2.588" transform="rotate(-42.481 21.66 35.525)"/><ellipse cx="36.295" cy="35.526" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.086" ry="2.588" transform="rotate(-42.481 36.295 35.526)"/>`),
		g.Group(children),
	)
}