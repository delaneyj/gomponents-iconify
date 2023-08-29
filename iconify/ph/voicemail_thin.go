package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 76a52 52 0 0 0-27.66 96H83.66A52 52 0 1 0 56 180h144a52 52 0 0 0 0-104ZM12 128a44 44 0 1 1 44 44a44.05 44.05 0 0 1-44-44Zm188 44a44 44 0 1 1 44-44a44.05 44.05 0 0 1-44 44Z"/>`),
		g.Group(children),
	)
}