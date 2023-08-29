package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M0 8v1600h1200V8H0zm194 865q-28 0-50-20.5T122 803t22-50t50-21q26 0 49 21.5t23 49.5t-22.5 49t-49.5 21z"/>`),
		g.Group(children),
	)
}