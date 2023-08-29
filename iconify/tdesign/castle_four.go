package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2h-8v6.342a5.961 5.961 0 0 0-1-.259V7h-2v1.083a5.962 5.962 0 0 0-1 .259V2H2v20h20V2ZM8 9.528A5.987 5.987 0 0 0 6 14v6H4V4h4v5.528ZM8 20v-5h8v5h-3v-3h-2v3H8Zm10 0v-6a5.987 5.987 0 0 0-2-4.472V4h4v16h-2Zm-2.126-7H8.126a4.003 4.003 0 0 1 7.748 0Z"/>`),
		g.Group(children),
	)
}