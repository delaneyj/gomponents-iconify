package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0zm9 0a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0z"/>`),
		g.Group(children),
	)
}