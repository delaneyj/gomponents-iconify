package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeanRemote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.311 13.452a13.949 13.949 0 0 1 4.697 10.507h0c-.002 7.545-5.825 13.66-13.008 13.66s-13.004-6.113-13.008-13.657a13.963 13.963 0 0 1 4.52-10.356M24 8.928v16.645"/>`),
		g.Group(children),
	)
}