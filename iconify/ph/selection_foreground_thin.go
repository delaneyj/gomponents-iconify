package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectionForegroundThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M60 216a4 4 0 0 1-4 4h-8a12 12 0 0 1-12-12v-8a4 4 0 0 1 8 0v8a4 4 0 0 0 4 4h8a4 4 0 0 1 4 4Zm52-4H96a4 4 0 0 0 0 8h16a4 4 0 0 0 0-8Zm-72-48a4 4 0 0 0 4-4v-16a4 4 0 0 0-8 0v16a4 4 0 0 0 4 4Zm128 32a4 4 0 0 0-4 4v8a4 4 0 0 1-4 4h-8a4 4 0 0 0 0 8h8a12 12 0 0 0 12-12v-8a4 4 0 0 0-4-4Zm0-88a4 4 0 0 0 4-4v-8a12 12 0 0 0-12-12h-8a4 4 0 0 0 0 8h8a4 4 0 0 1 4 4v8a4 4 0 0 0 4 4ZM56 84h-8a12 12 0 0 0-12 12v8a4 4 0 0 0 8 0v-8a4 4 0 0 1 4-4h8a4 4 0 0 0 0-8Zm152-48H96a12 12 0 0 0-12 12v40a4 4 0 0 0 4 4h24a4 4 0 0 0 0-8H92V48a4 4 0 0 1 4-4h112a4 4 0 0 1 4 4v112a4 4 0 0 1-4 4h-36v-20a4 4 0 0 0-8 0v24a4 4 0 0 0 3.51 3.95a2.17 2.17 0 0 0 .49.05h40a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}