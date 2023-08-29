package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBolderBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M177.1 114.5A48 48 0 0 0 140 36H64a12 12 0 0 0-12 12v152a12 12 0 0 0 12 12h88a52 52 0 0 0 25.1-97.5ZM76 60h64a24 24 0 0 1 0 48H76Zm76 128H76v-56h76a28 28 0 0 1 0 56Z"/>`),
		g.Group(children),
	)
}