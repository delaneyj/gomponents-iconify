package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M123.9 13.9H64.6l-6.4 10.5L100.9 51zM99.3 53.4L17.1 2.1v59.3l12.8 8zM74 102.7l36.9 23.2V66.6l-10.3-6.5zM28.7 74.6L4.1 114.1h59.3l34.7-55.5z"/>`),
		g.Group(children),
	)
}