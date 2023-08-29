package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M128.79 385q53 0 90.5 37.5t37.5 91t-37.5 91t-90.5 37.5t-90.5-37.5t-37.5-91t37.5-91t90.5-37.5zm842 212l-808 416q-38 20-81 9.5t-65.5-45t-11-72.5t50.5-58l646-333l-646-334q-39-19-50.5-57.5t11-73t65.5-45t81 9.5l808 416l2 2l2 1q23 13 36 33q22 34 10.5 72.5t-50.5 58.5z"/>`),
		g.Group(children),
	)
}