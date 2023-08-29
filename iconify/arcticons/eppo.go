package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eppo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.443 21.796c-1.09 1.636-1.09 3.272 0 4.909h2.728v5.454h2.181l2.182-4.909H25.08v3.273h3.818l1.636-3.273c17.455-.545 17.455-5.454 0-6l-1.636-3.273H25.08v3.273H12.534l-2.182-4.91H8.17v5.455H5.443Z"/>`),
		g.Group(children),
	)
}