package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M78 96a6 6 0 0 1-6 6H24a6 6 0 0 1 0-12h48a6 6 0 0 1 6 6Zm-6 26H24a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm0 32H24a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm0 32H24a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm80-64h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm0 32h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm0 32h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm80-96h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm-48-20h48a6 6 0 0 0 0-12h-48a6 6 0 0 0 0 12Zm48 52h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm0 32h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm0 32h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}