package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .001c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12zm0 1.614c4.997.005 9.168 3.533 10.164 8.234l.012.068H1.823C2.834 5.15 7.004 1.623 11.999 1.615zm0 20.772c-4.997-.005-9.169-3.533-10.165-8.234l-.012-.068h20.345c-1.003 4.769-5.173 8.298-10.168 8.301z"/>`),
		g.Group(children),
	)
}