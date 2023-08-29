package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 17.12c3.33-1.4 5.74-3.79 7.04-6.21c1.28-2.41 1.46-4.81.32-6.25c-1.03-1.29-2.37-1.78-3.73-1.74s-2.68.63-3.63 1.46c-.95-.83-2.27-1.42-3.63-1.46s-2.7.45-3.73 1.74c-1.14 1.44-.96 3.84.34 6.25c1.28 2.42 3.69 4.81 7.02 6.21z"/>`),
		g.Group(children),
	)
}