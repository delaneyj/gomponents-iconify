package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothFiletransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.82 34.212l20.048-19.933L20.075 4.5v39l9.793-9.779L9.82 13.788"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.177 18.789L24.936 24l5.241 5.211M24.936 24H38.18"/>`),
		g.Group(children),
	)
}