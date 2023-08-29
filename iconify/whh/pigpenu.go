package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M971 181L324 514l647 334q39 19 50.5 57.5t-11 73t-65.5 45t-82-9.5L55 597q-39-19-50.5-57.5T15 466q13-20 36-33q1 0 2-1t2-1L863 14q39-20 82-9.5t65.5 45t11 73T971 181z"/>`),
		g.Group(children),
	)
}