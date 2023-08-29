package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405.3 149.3C405.3 66.9 338.5 0 256 0c-82.5 0-149.3 66.9-149.3 149.3c0 51.7 26.3 97.3 66.3 124.1L106.7 512h298.7l-66.3-238.6c39.9-26.8 66.2-72.3 66.2-124.1z"/>`),
		g.Group(children),
	)
}