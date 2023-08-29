package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v12.414L14.414 22H2V2Zm2 2v16h9v-7h7V4H4Zm14.586 11H15v3.586L18.586 15ZM6 8h12v2H6V8Zm0 4h5v2H6v-2Z"/>`),
		g.Group(children),
	)
}