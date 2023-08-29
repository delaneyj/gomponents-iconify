package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaidan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.471 21.471 0 0 0 6.693 36.718a21.345 21.345 0 0 0-3.818 6.332a.995.995 0 0 0 .728 1.343C5.606 44.806 9.113 45.5 24 45.5a21.5 21.5 0 0 0 0-43Zm-9.5 12.851h19m-11 17.299h11m-19 0h4m-8-8.65h27"/>`),
		g.Group(children),
	)
}