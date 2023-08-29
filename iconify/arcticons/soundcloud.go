package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 30v3.26h14.42a5.08 5.08 0 0 0 0-10.16a5 5 0 0 0-1.56.27v-.27A8.3 8.3 0 0 0 24 16.17h0ZM4.5 24.34v6.46m2.78-8.81v10.46m2.78-9.7v10.48m2.78 0V19.9m2.78-1.01v14.34m2.78 0V19.9m2.78-1.94v15.27"/>`),
		g.Group(children),
	)
}