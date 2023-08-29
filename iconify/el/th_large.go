package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v525h525V0H0zm675 0v525h525V0H675zM0 675v525h525V675H0zm675 0v525h525V675H675z"/>`),
		g.Group(children),
	)
}