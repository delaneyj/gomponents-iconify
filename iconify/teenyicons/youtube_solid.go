package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8.599 7.5L7 8.566V6.434L8.599 7.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.506 1.773a28.606 28.606 0 0 1 11.988 0A1.905 1.905 0 0 1 15 3.636v7.728c0 .898-.628 1.675-1.506 1.863a28.605 28.605 0 0 1-11.988 0A1.905 1.905 0 0 1 0 11.364V3.636c0-.898.628-1.675 1.506-1.863Zm5.271 3.311A.5.5 0 0 0 6 5.5v4a.5.5 0 0 0 .777.416l3-2a.5.5 0 0 0 0-.832l-3-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}