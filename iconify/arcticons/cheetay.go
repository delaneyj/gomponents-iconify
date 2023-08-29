package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cheetay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.135 38.884a17.252 17.252 0 0 1-28.951-11.372"/><circle cx="27.39" cy="26.248" r="9.736" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.936" cy="8.601" r="4.101" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.414" cy="11.249" r="4.101" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.529" cy="11.8" r="4.101" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.586" cy="19.676" r="4.101" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}