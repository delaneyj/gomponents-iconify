package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 19.141v1.313c0 1.74-2.069 2.65-3.351 1.474l-7.04-6.454a2 2 0 0 1 0-2.948l7.04-6.454C11.93 4.896 14 5.806 14 7.546V8.86m3.04-2.787L10 12.526a2 2 0 0 0 0 2.948l7.04 6.454c1.283 1.176 3.352.266 3.352-1.475V17c5 0 8.5 10 8.5 10s1.5-16-8.5-16V7.546c0-1.74-2.069-2.65-3.352-1.474Z"/>`),
		g.Group(children),
	)
}