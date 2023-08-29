package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M32 0H4a4 4 0 0 0-4 4v28a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4z"/><path fill="#FFF" d="M19 7h-2a1 1 0 0 0-1 1v20a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1z"/><path fill="#FFF" d="m26.617 11.09l1.105 1.667a1 1 0 0 1-.281 1.386L10.769 25.191a1 1 0 0 1-1.386-.281l-1.105-1.667a1 1 0 0 1 .281-1.386L25.231 10.81a1 1 0 0 1 1.386.28z"/><path fill="#FFF" d="m9.383 11.09l-1.105 1.667a1 1 0 0 0 .281 1.386L25.231 25.19a1 1 0 0 0 1.386-.281l1.105-1.667a1 1 0 0 0-.281-1.386L10.769 10.809a1 1 0 0 0-1.386.281z"/>`),
		g.Group(children),
	)
}