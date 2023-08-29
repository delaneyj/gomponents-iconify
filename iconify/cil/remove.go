package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M10.655 17.03l3.97-3.97l3.97 3.97l1.061-1.061l-3.97-3.97l3.97-3.97l-1.061-1.061l-3.97 3.97l-3.97-3.97l-1.061 1.061l3.97 3.97l-3.97 3.97l1.061 1.061z" fill="currentColor"/><path d="M22.125 3H9.124a1.126 1.126 0 0 0-.816.351L.751 11.326v1.348l7.557 7.975c.206.216.495.351.816.351h13.001a1.127 1.127 0 0 0 1.125-1.125V4.125A1.127 1.127 0 0 0 22.125 3zm-.375 16.5H9.285L2.25 12.076v-.152L9.285 4.5H21.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}