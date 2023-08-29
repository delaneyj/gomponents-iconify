package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProstheticLeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 24v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1s4 1 4 1v8a2 2 0 0 0-2 2V19m-1.5 4.5H14m0 0h1.5m-1.5 0V20M11.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}