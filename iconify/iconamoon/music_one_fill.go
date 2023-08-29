package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.972 2A.996.996 0 0 0 11 3v11.535A4 4 0 1 0 13 18V8.18l4.836.806A1 1 0 0 0 19 8V4a1 1 0 0 0-.836-.986l-5.981-.997A1.004 1.004 0 0 0 11.972 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}