package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trailforks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1.608L0 22.392h24zm-1.292 5.698h2.584v5.885l2.664 1.917v5.587h-2.204V16.05L12 14.788l-1.752 1.262v4.645H8.044v-5.587l2.664-1.917z"/>`),
		g.Group(children),
	)
}