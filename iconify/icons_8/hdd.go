package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.22 6l-.19.75l-3 12l-.03.125V26h26v-7.125l-.03-.125l-3-12l-.19-.75H6.22zm1.56 2h16.44l2.5 10H5.28l2.5-10zM5 20h22v4H5v-4zm19 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}