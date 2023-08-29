package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PianoKeysFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM88 208H48V48h24v96a8 8 0 0 0 8 8h8Zm64 0h-48v-56h8a8 8 0 0 0 8-8V48h16v96a8 8 0 0 0 8 8h8Zm56 0h-40v-56h8a8 8 0 0 0 8-8V48h24v160Z"/>`),
		g.Group(children),
	)
}