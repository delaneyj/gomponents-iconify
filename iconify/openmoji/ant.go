package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#6A462F"><ellipse cx="36" cy="29.763" rx="6.263" ry="4.026"/><ellipse cx="36" cy="36.921" rx="4.474" ry="2.876"/><ellipse cx="36" cy="50.342" rx="6.263" ry="10.29"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m31.526 34.684l-8.052-3.579l-2.684-7.158m19.684 10.737l8.052-3.579l2.685-7.158M40.474 38.263h9.842L53 35.579m-21.474 2.684h-9.842L19 35.579m13.421 4.579l-9.842 8.947l.895 7.158m16.105-16.105l10.737 8.947l-.895 7.158M33.838 25.737l-1.417-3.579l.895-1.79m4.473 5.369l1.417-3.579l-.895-1.79"/><ellipse cx="36" cy="29.763" rx="6.263" ry="4.026"/><ellipse cx="36" cy="36.921" rx="4.474" ry="2.876"/><ellipse cx="36" cy="50.342" rx="6.263" ry="10.29"/></g>`),
		g.Group(children),
	)
}