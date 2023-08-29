package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.795 9.456l-9.5-6.923a.504.504 0 0 0-.59 0l-9.5 6.923a.502.502 0 0 0-.18.559l3.628 11.202a.5.5 0 0 0 .476.345H17.87a.5.5 0 0 0 .476-.345l3.629-11.202a.502.502 0 0 0-.181-.559zm-4.287 11.107H6.492L3.087 10.05L12 3.557l8.913 6.494l-3.405 10.511z"/>`),
		g.Group(children),
	)
}