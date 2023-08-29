package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.129 5.548a2.58 2.58 0 1 0 2.463 3.355H4.645a.774.774 0 0 1 0-1.548h3.17A4.13 4.13 0 0 1 16 8.129a4.13 4.13 0 0 1-8 1.44a4.13 4.13 0 0 1-8-1.44a4.129 4.129 0 0 1 6.353-3.48a.774.774 0 1 1-.835 1.305a2.565 2.565 0 0 0-1.389-.406Zm7.742 0a2.58 2.58 0 1 0 0 5.162a2.58 2.58 0 0 0 0-5.162Z"/>`),
		g.Group(children),
	)
}