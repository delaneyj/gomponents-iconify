package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12.9A5 5 0 1 0 11.098 9M15 12.9l-3.902-3.899l-7.513 8.584a2 2 0 1 0 2.827 2.83L15 12.9z"/>`),
		g.Group(children),
	)
}