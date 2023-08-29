package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VignetteLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42H40a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14Zm2 158a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2V56a2 2 0 0 1 2-2h176a2 2 0 0 1 2 2ZM128 74c-38.6 0-70 24.22-70 54s31.4 54 70 54s70-24.22 70-54s-31.4-54-70-54Zm0 96c-32 0-58-18.84-58-42s26-42 58-42s58 18.84 58 42s-26 42-58 42Z"/>`),
		g.Group(children),
	)
}