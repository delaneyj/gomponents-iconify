package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.987 29.856a19.51 19.51 0 0 1-1.784.63a19.249 19.249 0 0 1-2.39.572a8.714 8.714 0 0 1-1.971.145a5.262 5.262 0 0 1-2.114-.568a4.128 4.128 0 0 1-1.96-2.113a5.136 5.136 0 0 1-.11-2.923a11.848 11.848 0 0 1 1.85-4.098c.573-.847 1.218-1.642 1.872-2.428q.96-1.154 1.946-2.286a12.487 12.487 0 0 0-1.055 2.202a8.42 8.42 0 0 0-.476 1.91a4.3 4.3 0 0 0 .262 2.195a3.524 3.524 0 0 0 1.955 1.84a5.09 5.09 0 0 0 1.894.302a14.044 14.044 0 0 0 2.974-.441a82.76 82.76 0 0 0 3.254-.819C24.72 22.184 44.5 16.86 44.5 16.86s-22.854 9.82-30.513 12.996Z"/>`),
		g.Group(children),
	)
}