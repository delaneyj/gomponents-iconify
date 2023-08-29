package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M435.2 80H76.8c-24.9 0-44.6 19.6-44.6 44L32 388c0 24.4 19.9 44 44.8 44h358.4c24.9 0 44.8-19.6 44.8-44V124c0-24.4-19.9-44-44.8-44zm0 308H76.8V256h358.4v132zm0-220H76.8v-44h358.4v44z" fill="currentColor"/>`),
		g.Group(children),
	)
}