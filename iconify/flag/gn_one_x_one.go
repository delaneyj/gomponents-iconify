package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="red" d="M0 0h170.7v512H0z"/><path fill="#ff0" d="M170.7 0h170.6v512H170.7z"/><path fill="#090" d="M341.3 0H512v512H341.3z"/></g>`),
		g.Group(children),
	)
}