package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.762 3.25c.46.006.99.223 1.276.734l.174.312a82.518 82.518 0 0 1 6.474 14.753c.218.67-.14 1.237-.589 1.5a1.484 1.484 0 0 1-1.552-.037l-4.728-3.04c-.522-.336-1.231-.337-1.72-.028L6.464 20.37a1.483 1.483 0 0 1-1.553.01c-.427-.255-.809-.808-.61-1.482A77.671 77.671 0 0 1 10.307 4.29l.164-.308c.284-.531.832-.737 1.29-.73Zm-.007 1.513l-.125.233A76.203 76.203 0 0 0 5.84 18.99l4.455-2.813c1-.63 2.337-.607 3.334.034l4.496 2.892a81.088 81.088 0 0 0-6.222-14.074l-.148-.265Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}