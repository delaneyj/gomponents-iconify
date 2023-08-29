package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M169.53 16.344L259.345 88L337 92.28l-1.03 18.657l-161.376-8.906l-118.78-4.905l227.28 68.03l-197.72 246.75l-14.53-17.655l-49.22 96.625l248.69-202.78l51.81 11.592l-38.78 40.594L270.5 329.5l-57.28 84.125L444.843 273.47L328 241.06l100.22-81.718c1.132.46 2.3.898 3.5 1.22c23.324 6.248 49.764-16.835 59.06-51.533c9.298-34.695-2.08-67.874-25.405-74.124c-23.325-6.25-49.765 16.802-59.063 51.5a95.43 95.43 0 0 0-2.875 16.22L169.53 16.343z"/>`),
		g.Group(children),
	)
}