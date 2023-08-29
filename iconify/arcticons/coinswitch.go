package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coinswitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.169" cy="8.393" r="3.893" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.832" cy="38.682" r="3.893" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.397 6.326c-4.093 0-11.133 2.237-11.133 8.295c0 11.734 21.993 8.077 21.993 18.501c0 7.913-7.695 11.752-14.498 9.932"/>`),
		g.Group(children),
	)
}