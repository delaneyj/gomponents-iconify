package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.982 11.19l4.177-4.185a3.604 3.604 0 0 0 3.795-.832c.698-.701 1.027-1.534 1.084-2.506l-.677.694l-2.212.74l-2.269-2.257l.776-2.174l.672-.67c-.946.078-1.799.36-2.487 1.048a3.623 3.623 0 0 0-.84 3.772L5.852 8.978c-1.266-.456-2.737-.117-3.765.914a3.596 3.596 0 0 0 0 5.078c1.387 1.39 3.654 1.37 5.066-.045c1.002-1.003 1.231-2.491.83-3.736zm-5.25 3.1a2.5 2.5 0 1 1 3.535-3.535a2.5 2.5 0 0 1-3.534 3.535z"/>`),
		g.Group(children),
	)
}