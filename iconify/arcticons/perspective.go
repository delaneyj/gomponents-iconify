package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Perspective(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.184 14.341a3.226 3.226 0 1 1-6.452 0a3.226 3.226 0 0 1 6.452 0Zm-2.232 17.157a3.226 3.226 0 1 1-6.452 0a3.226 3.226 0 0 1 6.452 0ZM43.5 38.077a3.226 3.226 0 0 1-3.226 3.226h0a3.226 3.226 0 0 1-3.226-3.226h0a3.226 3.226 0 0 1 3.226-3.226h0a3.226 3.226 0 0 1 3.226 3.226h0ZM41.378 9.923a3.226 3.226 0 0 1-3.226 3.226h0a3.226 3.226 0 0 1-3.226-3.226h0a3.226 3.226 0 0 1 3.226-3.226h0a3.226 3.226 0 0 1 3.226 3.226h0ZM13.1 13.82l21.876-3.562m3.167 2.89l1.847 21.75m-2.849 2.591l-26.32-5.164m-2.685-3.989l1.486-10.823"/>`),
		g.Group(children),
	)
}