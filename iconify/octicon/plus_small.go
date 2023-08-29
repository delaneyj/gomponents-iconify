package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M4 4H3v3H0v1h3v3h1V8h3V7H4V4z" fill="currentColor"/>`),
		g.Group(children),
	)
}