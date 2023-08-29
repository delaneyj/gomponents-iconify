package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 281.073V0h192v94.65C106.01 132.14 30.952 197.93 0 281.072zM0 416v96h192v-96c0-160 256-160 256-160V64S0 64 0 416z"/>`),
		g.Group(children),
	)
}