package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyAngryBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 156a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm72-32a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm72 4A108 108 0 1 1 128 20a108.12 108.12 0 0 1 108 108Zm-24 0a84 84 0 1 0-84 84a84.09 84.09 0 0 0 84-84ZM85.34 102l36 24a12 12 0 0 0 13.32 0l36-24a12 12 0 0 0-13.32-20L128 101.58L98.66 82a12 12 0 0 0-13.32 20ZM154 167.12a51.1 51.1 0 0 0-52 0a12 12 0 1 0 12 20.76a27.13 27.13 0 0 1 28 0a12 12 0 1 0 12-20.76Z"/>`),
		g.Group(children),
	)
}