package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.501 11.65a5.25 5.25 0 1 1-10.5 0a5.25 5.25 0 0 1 10.5 0zm13.5 0a5.25 5.25 0 1 1-10.5 0a5.25 5.25 0 0 1 10.5 0z"/>`),
		g.Group(children),
	)
}