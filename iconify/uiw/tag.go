package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.368 0a.68.68 0 0 1 .484.201l8.35 8.39c.558.586.836 1.249.794 1.958c-.04.664-.304 1.274-.807 1.846l-7.074 7.037l-.115.092c-.638.406-1.26.564-1.85.428c-.515-.118-1.054-.436-1.677-.978l-8.16-8.219a.68.68 0 0 1-.199-.47L0 1.472C.007 1.044.126.681.392.413C.666.138 1.055.023 1.588 0h8.78ZM6.473 4.574a1.59 1.59 0 1 0 0 3.18a1.59 1.59 0 1 0 0-3.18Z"/>`),
		g.Group(children),
	)
}