package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdColorWand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M200.8 157.2l-36.4 37.4L411.7 448l36.3-37.4z" fill="currentColor"/><path d="M181 64h37v68h-37z" fill="currentColor"/><path d="M181 262h37v68h-37z" fill="currentColor"/><path d="M270 176h69v37h-69z" fill="currentColor"/><path d="M305.6 115.8l-25.7-26.3-47.1 48.3 25.6 26.2z" fill="currentColor"/><path d="M168.8 137.8l-47.1-48.3-25.6 26.3 47.1 48.2z" fill="currentColor"/><path d="M96.1 277.9l25.6 26.2 47.1-48.2-25.6-26.3z" fill="currentColor"/><path d="M64 176h65v37H64z" fill="currentColor"/>`),
		g.Group(children),
	)
}