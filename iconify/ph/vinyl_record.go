package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm0-144a56.06 56.06 0 0 0-56 56a8 8 0 0 1-16 0a72.08 72.08 0 0 1 72-72a8 8 0 0 1 0 16Zm72 56a72.08 72.08 0 0 1-72 72a8 8 0 0 1 0-16a56.06 56.06 0 0 0 56-56a8 8 0 0 1 16 0Zm-40 0a32 32 0 1 0-32 32a32 32 0 0 0 32-32Zm-48 0a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}