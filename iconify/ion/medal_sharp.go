package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M80 32L16 160h289.11l80.22-128H80z"/><path fill="currentColor" d="M496 144L424 32L298 231.08a128 128 0 0 0-84 0L189.53 192H43.82l86.66 134.89a128 128 0 1 0 251 0ZM256 422a70 70 0 1 1 70-70a70.08 70.08 0 0 1-70 70Z"/><circle cx="256" cy="352" r="32" fill="currentColor"/>`),
		g.Group(children),
	)
}