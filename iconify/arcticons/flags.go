package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 28.604c13.168 7.685 25.352-7.97 37.999-.096M5 19.587c13.168 7.684 25.353-7.972 38-.096"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 10.569c13.168 7.684 25.353-7.972 38-.096h-.001v27.053h0C30.353 29.651 18.168 45.307 5 37.622h0V10.57h0Z"/>`),
		g.Group(children),
	)
}