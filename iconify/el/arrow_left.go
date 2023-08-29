package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m0 600l703.125 468.75V816.211H1200V383.789H703.125V131.25L0 600z"/>`),
		g.Group(children),
	)
}