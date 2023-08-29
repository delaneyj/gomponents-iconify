package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownSquareF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm2.038 13.23h1.005l1.823 1.71a2 2 0 0 0 1.367.54h.205a1.6 1.6 0 0 0 1.6-1.6v-7.8a1.6 1.6 0 0 0-1.6-1.6h-.205a2 2 0 0 0-1.367.541l-1.823 1.71H6.038a2 2 0 0 0-2 2v2.5a2 2 0 0 0 2 2zm7.664-4.24a1 1 0 0 1 0 2a1 1 0 1 0 0 2a3 3 0 0 0 0-6a1 1 0 1 0 0 2zm-5.869-.26l2.205-2.066v6.633L7.833 11.23H6.038v-2.5h1.795z"/>`),
		g.Group(children),
	)
}