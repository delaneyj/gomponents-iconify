package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M198.114 256h115.773v200.186H198.114z"/><path fill="currentColor" d="M0 0v512h512V0H0zm484.346 484.346H27.654V256h170.46V55.814h115.772V256h170.46v228.346z"/>`),
		g.Group(children),
	)
}