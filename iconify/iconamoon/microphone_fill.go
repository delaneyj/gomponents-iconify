package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a4 4 0 0 0-4 4v5a4 4 0 0 0 8 0V6a4 4 0 0 0-4-4Zm-7 8a1 1 0 0 1 1 1a6 6 0 0 0 12 0a1 1 0 1 1 2 0a8.001 8.001 0 0 1-7 7.938V21a1 1 0 1 1-2 0v-2.062A8.001 8.001 0 0 1 4 11a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}