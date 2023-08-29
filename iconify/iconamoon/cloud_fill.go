package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.303 6.697a7.5 7.5 0 0 0-12.47 3.088A5.002 5.002 0 0 0 6.5 19.5h12a4 4 0 0 0 .99-7.876a7.474 7.474 0 0 0-2.187-4.927Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}