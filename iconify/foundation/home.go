package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M83.505 37.85L51.013 12.688a2.508 2.508 0 0 0-3.1.025L16.46 37.874a2.509 2.509 0 0 0-.939 1.956v45.5a2.504 2.504 0 0 0 2.505 2.505h18.697a2.506 2.506 0 0 0 2.505-2.505V57.471h21.54V85.33a2.505 2.505 0 0 0 2.505 2.505h18.7a2.506 2.506 0 0 0 2.505-2.505v-45.5a2.498 2.498 0 0 0-.973-1.98z"/>`),
		g.Group(children),
	)
}