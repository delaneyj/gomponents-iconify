package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2a3 3 0 0 0-3 3v6a3 3 0 0 0 6 0V5a3 3 0 0 0-3-3zm0-2a5 5 0 0 1 5 5v6a5 5 0 0 1-10 0V5a5 5 0 0 1 5-5zM0 11.03a1 1 0 1 1 2 0A6.97 6.97 0 0 0 8.97 18h.06A6.97 6.97 0 0 0 16 11.03a1 1 0 1 1 2 0A8.97 8.97 0 0 1 9.03 20h-.06A8.97 8.97 0 0 1 0 11.03z"/>`),
		g.Group(children),
	)
}