package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m972.38 354l-295-280q-22-23-71-48.5t-79-25.5h-376q-63 0-107 43.5T.38 149v598q0 62 44 105.5t107 43.5h376q30 0 82-25t68-50l295-279q52-40 52-94t-52-94z"/>`),
		g.Group(children),
	)
}