package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.432 4.079a1.25 1.25 0 0 0-2.033.391l-1.76 4.105a4.25 4.25 0 0 0-4.215 1.07L7.067 11a.75.75 0 0 0 0 1.06l2.142 2.143l-5.74 5.739a.75.75 0 1 0 1.061 1.06l5.74-5.739l2.141 2.142a.75.75 0 0 0 1.06 0l1.357-1.356a4.25 4.25 0 0 0 1.07-4.217l4.105-1.76a1.25 1.25 0 0 0 .392-2.032l-3.963-3.96Z"/>`),
		g.Group(children),
	)
}