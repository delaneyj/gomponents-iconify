package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronWithCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 10L8.698 7.494a.512.512 0 0 1 0-.718a.5.5 0 0 1 .71 0l2.807 2.864a.51.51 0 0 1 0 .717l-2.807 2.864a.498.498 0 0 1-.71 0a.51.51 0 0 1 0-.717L11 10zM10 .4a9.6 9.6 0 0 1 9.6 9.6c0 5.303-4.298 9.6-9.6 9.6S.4 15.303.4 10A9.6 9.6 0 0 1 10 .4zm0 17.954a8.354 8.354 0 1 0 0-16.709a8.354 8.354 0 0 0 0 16.709z"/>`),
		g.Group(children),
	)
}