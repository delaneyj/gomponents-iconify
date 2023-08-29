package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 512H397.47v-56.587h52.667V56.86H397.47V0H512zM375.885 107.743v59.344h-91.161v275.88h-64.107v-275.88h-91.586v-59.344zM0 0h113.963v56.587H61.295V455.14h52.668V512H0z"/>`),
		g.Group(children),
	)
}