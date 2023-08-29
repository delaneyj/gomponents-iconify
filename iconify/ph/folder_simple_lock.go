package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSimpleLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 88v16a8 8 0 0 1-16 0V88h-85.33a16.12 16.12 0 0 1-9.6-3.2L93.33 64H40v136h72a8 8 0 0 1 0 16H40a16 16 0 0 1-16-16V64a16 16 0 0 1 16-16h53.33a16.12 16.12 0 0 1 9.6 3.2L130.67 72H216a16 16 0 0 1 16 16Zm0 80v40a8 8 0 0 1-8 8h-72a8 8 0 0 1-8-8v-40a8 8 0 0 1 8-8h8v-4a28 28 0 0 1 56 0v4h8a8 8 0 0 1 8 8Zm-56-8h24v-4a12 12 0 0 0-24 0Zm40 16h-56v24h56Z"/>`),
		g.Group(children),
	)
}