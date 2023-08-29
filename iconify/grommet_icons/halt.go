package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Halt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 23h7c3 0 4-2 4-5V6c0-2-1-2-1.5-2S16 4 16 6v7v-9c0-1 0-2-1.5-2S13 3 13 4v9V3c0-1 0-2-1.5-2S10 2 10 3v10v-9c0-1 .03-2-1.5-2C7 2 7 3 7 4v14v-4c0-1-.55-2-2-2H4v6c0 3.962 2 5.024 4 5Z"/>`),
		g.Group(children),
	)
}