package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mooltifill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="29.135" height="25.056" x="9.432" y="18.444" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><circle cx="24" cy="30.973" r="3.496" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.26 13.24A8.74 8.74 0 0 1 24 4.5h0a8.74 8.74 0 0 1 8.74 8.74m0 0v5.204M15.26 13.24l.01 5.204m-.01-5.204v5.204"/>`),
		g.Group(children),
	)
}