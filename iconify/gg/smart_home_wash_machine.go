package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartHomeWashMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 4h12a1 1 0 0 1 1 1v3H5V5a1 1 0 0 1 1-1Zm13 15v-9H5v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1ZM3 5a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V5Zm4 0a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2H7Zm7 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-4 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm2 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}