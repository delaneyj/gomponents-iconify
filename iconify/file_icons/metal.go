package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 464.102V47.898l194.972 236.186V109.719L512 464.102h-79.257L255.207 269.818v194.284L64.374 225.15v238.952z"/>`),
		g.Group(children),
	)
}