package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0ZM5.252 4.7a1.5 1.5 0 0 1 2.047.55a.5.5 0 0 1-.866.5a.5.5 0 0 0-.865-.001a.5.5 0 1 1-.865-.502a1.5 1.5 0 0 1 .55-.547ZM8 12a4 4 0 0 0 4-4H4a4 4 0 0 0 4 4Zm1.252-7.3a1.5 1.5 0 0 1 2.047.55a.5.5 0 1 1-.866.5a.5.5 0 0 0-.865-.001a.5.5 0 1 1-.865-.502a1.5 1.5 0 0 1 .55-.547Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}