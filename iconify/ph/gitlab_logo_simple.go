package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitlabLogoSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230.15 117.1L210.25 41a11.94 11.94 0 0 0-22.79-1.11L169.78 88H86.22L68.54 39.87A11.94 11.94 0 0 0 45.75 41l-19.9 76.1a57.19 57.19 0 0 0 22 61l73.27 51.76a11.91 11.91 0 0 0 13.74 0l73.27-51.76a57.19 57.19 0 0 0 22.02-61ZM198.91 165L128 215.09L57.09 165a41.1 41.1 0 0 1-15.75-43.84L58 57.5l15.13 41.26a8 8 0 0 0 7.51 5.24h94.72a8 8 0 0 0 7.51-5.24L198 57.5l16.63 63.65A41.1 41.1 0 0 1 198.91 165Z"/>`),
		g.Group(children),
	)
}