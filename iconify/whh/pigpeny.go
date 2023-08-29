package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpeny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M897 642q-53 0-90.5-37.5t-37.5-91t37.5-91T897 385t90.5 37.5t37.5 91t-37.5 91T897 642zm73-462L324 514l646 333q39 20 50.5 58t-11 72.5t-65.5 45t-81-9.5L55 597q-39-20-50.5-58.5T15 466q13-20 36-33q1 0 2-1l2-2L863 14q38-20 81-9.5t65.5 45t11 73T970 180z"/>`),
		g.Group(children),
	)
}