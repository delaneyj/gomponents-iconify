package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M10 7H6l3-7l-9 9h4l-3 7l9-9z" fill="currentColor"/>`),
		g.Group(children),
	)
}