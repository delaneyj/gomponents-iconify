package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparklingHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ff5a79" d="M61.1 18.2c-6.4-17-27.2-9.4-29.1-.9c-2.6-9-22.9-15.7-29.1.9C-4 36.7 29.6 53.3 32 56c2.4-2.2 36-19.6 29.1-37.8"/><path fill="#fff" d="M53.8 13.3L51.2 8l-2.6 5.3l-5.3 2.6l5.3 2.6l2.6 5.3l2.6-5.3l5.3-2.6zM16.3 26.8l-3.6-7.2l-3.6 7.2L2 30.3l7.1 3.6l3.6 7.1l3.6-7.1l7.1-3.6zm32.4 15.6l-2.8-5.7l-2.9 5.7l-5.8 2.9l5.8 2.9l2.9 5.7l2.8-5.7l5.8-2.9zM18.886 20.601l2.97-2.97l2.97 2.97l-2.97 2.97zm4.491 20.973l2.475-2.475l2.474 2.475l-2.474 2.474zm15.937-14.282l2.97-2.97l2.97 2.97l-2.97 2.97z"/>`),
		g.Group(children),
	)
}