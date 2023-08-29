package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.13 9h9.75v1.25H3.13zM.5 5.75h15V7H.5zm0 6.5h15v1.25H.5zM3.13 2.5h9.75v1.25H3.13z"/>`),
		g.Group(children),
	)
}