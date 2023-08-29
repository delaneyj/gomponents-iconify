package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 96a8 8 0 0 1-8 8H24a8 8 0 0 1 0-16h48a8 8 0 0 1 8 8Zm-8 24H24a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm0 32H24a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm0 32H24a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm80-64h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm0 32h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm0 32h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm80-96h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm-48-16h48a8 8 0 0 0 0-16h-48a8 8 0 0 0 0 16Zm48 48h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm0 32h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Zm0 32h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}