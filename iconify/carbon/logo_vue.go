package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoVue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.9 3.9L16 19.4L7.1 3.9H2l14 24.3L30 3.9h-5.1z"/><path fill="currentColor" d="m16 9.5l-3.2-5.6H8.1L16 17.5l7.9-13.6h-4.6L16 9.5z"/>`),
		g.Group(children),
	)
}