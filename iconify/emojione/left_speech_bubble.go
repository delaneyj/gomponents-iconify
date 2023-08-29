package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4fd1d9" d="M8.6 5h46.9c3.6 0 6.6 2.9 6.6 6.5v29.8c0 3.6-2.9 6.5-6.6 6.5h-6.9V59L38.1 47.8H8.6C5 47.8 2 44.9 2 41.3V11.5C2 7.9 4.9 5 8.6 5z"/><g fill="#fff"><circle cx="48" cy="26.4" r="4.2"/><circle cx="32" cy="26.4" r="4.2"/><circle cx="16" cy="26.4" r="4.2"/></g>`),
		g.Group(children),
	)
}