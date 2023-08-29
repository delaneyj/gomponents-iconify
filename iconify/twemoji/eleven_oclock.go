package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevenOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#99AAB5"/><circle cx="18" cy="18" r="14" fill="#E1E8ED"/><path fill="#67757F" d="M18 19a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v12a1 1 0 0 1-1 1z"/><path fill="#67757F" d="M18.001 19a1 1 0 0 1-.896-.553l-4-8a1 1 0 0 1 1.789-.895l4 8A1 1 0 0 1 18.001 19z"/>`),
		g.Group(children),
	)
}