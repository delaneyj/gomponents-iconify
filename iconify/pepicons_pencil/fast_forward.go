package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.5 14.796L8.981 10L3.5 5.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506L4.159 4.451c-.647-.566-1.659-.106-1.659.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z"/><path d="M10.5 14.796L15.981 10L10.5 5.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506l-5.482-4.796c-.646-.566-1.658-.106-1.658.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z"/></g>`),
		g.Group(children),
	)
}