package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swissid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="4.694" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="10.194" r="4.694" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="37.806" r="4.694" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.806" cy="24" r="4.694" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="10.194" cy="24" r="4.694" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}