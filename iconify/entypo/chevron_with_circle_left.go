package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronWithCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.302 6.776a.5.5 0 0 0-.71 0L7.785 9.641a.51.51 0 0 0 0 .717l2.807 2.864a.498.498 0 0 0 .71 0a.51.51 0 0 0 0-.717L9 10l2.302-2.506a.512.512 0 0 0 0-.718zM10 .4A9.6 9.6 0 0 0 .4 10c0 5.303 4.298 9.6 9.6 9.6s9.6-4.297 9.6-9.6A9.6 9.6 0 0 0 10 .4zm0 17.954A8.353 8.353 0 0 1 1.646 10A8.354 8.354 0 1 1 10 18.354z"/>`),
		g.Group(children),
	)
}