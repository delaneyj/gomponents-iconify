package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DmUnipd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h15.725v15.725H5.5M42.5 5.5v15.725H26.775V5.5m-5.55 37V26.775H5.5V42.5m37-15.725H26.775V42.5H42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5a7.862 7.862 0 0 1 0 15.725h0V5.5Zm37 0a7.862 7.862 0 0 1-15.725 0H42.5Zm-37 37a7.862 7.862 0 0 1 15.725 0H5.5Zm37 0a7.862 7.862 0 0 1 0-15.725V42.5Z"/>`),
		g.Group(children),
	)
}