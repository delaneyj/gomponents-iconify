package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h24V0Zm5.136 7.02c1.416 0 3.162 1.158 3.288 3.072c0 0 .222 3.81.228 4.11c.12-.318 1.602-4.146 1.602-4.146c.588-1.56 1.692-3.036 3.768-3.036c0 0 .888 6.858.924 7.176c.09-.324 1.308-4.14 1.308-4.14c.582-1.566 1.716-3.036 3.87-3.036l-3.054 7.506c-.63 1.506-1.806 2.88-3.984 2.88c0 0-.924-6.426-.936-6.57c-.054.15-1.356 3.528-1.356 3.528c-.612 1.566-1.782 3.036-3.954 3.042z"/>`),
		g.Group(children),
	)
}