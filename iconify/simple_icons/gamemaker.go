package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamemaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.012 11.994L12.006 0l11.982 12.006h-6.831l-5.163-5.151l-5.151 5.151l5.163 5.151v-5.151h5.151v6.903L12.006 24z"/>`),
		g.Group(children),
	)
}