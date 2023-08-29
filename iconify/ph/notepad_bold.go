package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 124a12 12 0 0 1-12 12H96a12 12 0 0 1 0-24h64a12 12 0 0 1 12 12Zm-12 28H96a12 12 0 0 0 0 24h64a12 12 0 0 0 0-24Zm60-104v152a36 36 0 0 1-36 36H72a36 36 0 0 1-36-36V48a20 20 0 0 1 20-20h16v-4a12 12 0 0 1 24 0v4h20v-4a12 12 0 0 1 24 0v4h20v-4a12 12 0 0 1 24 0v4h16a20 20 0 0 1 20 20Zm-24 4h-12v4a12 12 0 0 1-24 0v-4h-20v4a12 12 0 0 1-24 0v-4H96v4a12 12 0 0 1-24 0v-4H60v148a12 12 0 0 0 12 12h112a12 12 0 0 0 12-12Z"/>`),
		g.Group(children),
	)
}