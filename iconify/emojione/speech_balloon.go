package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4fd1d9" d="M55.4 5H8.6C4.9 5 2 7.9 2 11.5v29.8c0 3.6 2.9 6.5 6.6 6.5h6.9V59l10.4-11.2h29.6c3.6 0 6.6-2.9 6.6-6.5V11.5C62 7.9 59.1 5 55.4 5z"/><g fill="#fff"><circle cx="16" cy="26.4" r="4.2"/><circle cx="32" cy="26.4" r="4.2"/><circle cx="48" cy="26.4" r="4.2"/></g>`),
		g.Group(children),
	)
}