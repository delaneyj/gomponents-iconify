package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicRecordTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 6.5a4.5 4.5 0 1 1 9 0v7.124a7.513 7.513 0 0 0-4.35 5.373L13 19a4.5 4.5 0 0 1-4.5-4.5v-8ZM13.016 21H13a6.5 6.5 0 0 1-6.5-6.5v-.75a.75.75 0 1 0-1.5 0v.75a8 8 0 0 0 7.25 7.965v2.785a.75.75 0 0 0 1.5 0v-1.477A7.457 7.457 0 0 1 13.016 21Zm7.484 4.5a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0 1.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Zm0-3a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"/>`),
		g.Group(children),
	)
}