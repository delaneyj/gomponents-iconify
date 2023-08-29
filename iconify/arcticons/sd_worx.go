package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdWorx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.644 4.5l-6.547 31.05m-8.511 7.95l2.993-26.935M12.356 24.14l1.871 11.036"/>`),
		g.Group(children),
	)
}