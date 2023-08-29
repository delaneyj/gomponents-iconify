package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductivityLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 30.802c2.28-2.64 5.353-5.374 8.25-4.768c3.702.775 4.508 6.462 6.91 6.322c3.408-.2 4.453-11.795 8.197-11.947c1.436-.058 1.931 1.621 3.643 1.661c2.496.058 4.603-3.44 6-6.429"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}