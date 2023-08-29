package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 14c-2.188 0-3.96-1.789-3.96-4c0-2.211 1.772-4 3.96-4c2.187 0 3.96 1.789 3.96 4c0 2.211-1.773 4-3.96 4zm10 0c-2.188 0-3.96-1.789-3.96-4c0-2.211 1.772-4 3.96-4c2.187 0 3.96 1.789 3.96 4c0 2.211-1.773 4-3.96 4z"/>`),
		g.Group(children),
	)
}