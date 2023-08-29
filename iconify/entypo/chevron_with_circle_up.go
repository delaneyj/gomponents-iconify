package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronWithCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.359 7.785a.51.51 0 0 0-.717 0l-2.864 2.807a.498.498 0 0 0 0 .71a.51.51 0 0 0 .717 0L10 9l2.506 2.302a.512.512 0 0 0 .718 0a.5.5 0 0 0 0-.71l-2.865-2.807zM10 .4A9.6 9.6 0 0 0 .4 10c0 5.303 4.298 9.6 9.6 9.6s9.6-4.297 9.6-9.6A9.6 9.6 0 0 0 10 .4zm0 17.954A8.353 8.353 0 0 1 1.646 10A8.354 8.354 0 1 1 10 18.354z"/>`),
		g.Group(children),
	)
}