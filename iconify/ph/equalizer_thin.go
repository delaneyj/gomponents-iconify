package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M76 96a4 4 0 0 1-4 4H24a4 4 0 0 1 0-8h48a4 4 0 0 1 4 4Zm-4 28H24a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm0 32H24a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm0 32H24a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm80-64h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm0 32h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm0 32h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm80-96h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm-48-24h48a4 4 0 0 0 0-8h-48a4 4 0 0 0 0 8Zm48 56h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm0 32h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Zm0 32h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}