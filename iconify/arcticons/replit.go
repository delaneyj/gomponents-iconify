package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Replit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="6.295" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="10.863" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.496" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.127 29.041a16.062 16.062 0 0 1 24.018-15.438"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.959 3.127a16.062 16.062 0 0 1 15.438 24.018"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.872 18.959a16.062 16.062 0 0 1-24.018 15.438"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.04 44.872a16.062 16.062 0 0 1-15.438-24.018"/>`),
		g.Group(children),
	)
}