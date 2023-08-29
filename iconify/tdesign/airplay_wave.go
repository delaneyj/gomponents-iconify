package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v17h-5.5v-2H21V5H3l.001 13h3.5v2H1V3Zm15.95 10.383a7 7 0 0 0-9.9 0l-.706.707l-1.414-1.414l.707-.707a9 9 0 0 1 12.728 0l.707.707l-1.415 1.414l-.707-.707Zm-2.828 2.828a3 3 0 0 0-4.243 0l-.707.707l-1.414-1.414l.707-.707a5 5 0 0 1 7.07 0l.708.707l-1.414 1.414l-.707-.707ZM12 18.086L15.914 22H8.086L12 18.086Z"/>`),
		g.Group(children),
	)
}