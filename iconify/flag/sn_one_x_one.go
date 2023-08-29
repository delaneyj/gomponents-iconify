package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#0b7226" d="M0 0h170.7v512H0z"/><path fill="#ff0" d="M170.7 0h170.6v512H170.7z"/><path fill="#bc0000" d="M341.3 0H512v512H341.3z"/></g><path fill="#0b7226" d="m197 351.7l22-71.7l-60.4-46.5h74.5l24.2-76l22.1 76H356L295.6 280l22.1 74l-60.3-46.5z"/>`),
		g.Group(children),
	)
}