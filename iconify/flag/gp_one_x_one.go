package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0z"/><path fill="#002654" d="M0 0h170.7v512H0z"/><path fill="#ce1126" d="M341.3 0H512v512H341.3z"/>`),
		g.Group(children),
	)
}