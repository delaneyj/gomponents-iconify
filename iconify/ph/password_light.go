package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M46 56v144a6 6 0 0 1-12 0V56a6 6 0 0 1 12 0Zm86.58 56.41L110 119.74V96a6 6 0 0 0-12 0v23.74l-22.58-7.33a6 6 0 1 0-3.71 11.41l22.58 7.33l-14 19.21a6 6 0 1 0 9.7 7.06l14-19.21l14 19.21a6 6 0 0 0 9.7-7.06l-14-19.21l22.58-7.33a6 6 0 1 0-3.71-11.41Zm103.56 3.85a6 6 0 0 0-7.56-3.85L206 119.74V96a6 6 0 0 0-12 0v23.74l-22.58-7.33a6 6 0 1 0-3.71 11.41l22.58 7.33l-13.95 19.21a6 6 0 1 0 9.7 7.06l14-19.21l14 19.21a6 6 0 0 0 9.7-7.06l-13.95-19.21l22.58-7.33a6 6 0 0 0 3.77-7.56Z"/>`),
		g.Group(children),
	)
}