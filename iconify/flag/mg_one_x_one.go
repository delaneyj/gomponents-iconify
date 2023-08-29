package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MgOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#fc3d32" d="M170.7 0H512v256H170.7z"/><path fill="#007e3a" d="M170.7 256H512v256H170.7z"/><path fill="#fff" d="M0 0h170.7v512H0z"/></g>`),
		g.Group(children),
	)
}