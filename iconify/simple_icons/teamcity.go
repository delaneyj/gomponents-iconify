package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teamcity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h24V0zm2.664 2.964h7.48v1.832H7.396v7.196H5.412V4.796H2.664zm9.328 18h-9v-1.5h9zm5.564-9.218a4.62 4.62 0 0 1-2.036.374a4.556 4.556 0 0 1-4.628-4.616V7.48A4.584 4.584 0 0 1 15.6 2.812A4.656 4.656 0 0 1 19.16 4.2l-1.264 1.456a3.336 3.336 0 0 0-2.312-1.02a2.671 2.671 0 0 0-2.616 2.8v.028a2.68 2.68 0 0 0 2.616 2.836a3.226 3.226 0 0 0 2.376-1.056l1.264 1.276a4.62 4.62 0 0 1-1.668 1.226z"/>`),
		g.Group(children),
	)
}