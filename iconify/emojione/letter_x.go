package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M26.2 46.5h-7l9.3-14.8l-8.9-14.2h7.2l5.2 9.4l5.3-9.4h7l-8.9 14l9.4 15h-7.4L32 36.6l-5.8 9.9"/>`),
		g.Group(children),
	)
}