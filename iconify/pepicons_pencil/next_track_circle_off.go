package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 17.796L11.981 13L6.5 8.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506L7.159 7.451c-.647-.566-1.659-.106-1.659.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.5 17.796L18.981 13L13.5 8.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506l-5.482-4.796c-.646-.566-1.658-.106-1.658.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 7.5a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-1 0V8a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}