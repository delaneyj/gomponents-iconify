package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Curve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.495 43.295a21.5 21.5 0 1 1 .368-38.405m3.561 3.957L21.66 20.034m21.56-3.458L27.153 28.068m17.694-1.628l-12.203 9.255"/>`),
		g.Group(children),
	)
}