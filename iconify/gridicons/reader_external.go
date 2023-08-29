package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderExternal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.376 2.016h6.608v6.608h-1.5V4.577l-5.87 5.87l-.53.53l-1.061-1.06l.53-.53l5.87-5.87h-4.047zM5 5.5h4V4H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-4h-1.5v4a.5.5 0 0 1-.5.5H5a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 .5-.5z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}