package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h12.515L21 9.68V22H3V2Zm9 2h-1v3.5H9V4H8v3.5H6V4H5v16h14v-9.68L14.485 4H14v3.5h-2V4Z"/>`),
		g.Group(children),
	)
}