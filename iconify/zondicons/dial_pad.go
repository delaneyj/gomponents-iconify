package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialPad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm5 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm5 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zM5 9a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm5 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm5 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zM5 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm5 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm5-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}