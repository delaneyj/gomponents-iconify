package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilNextTrackCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.5 17.796L11.981 13L6.5 8.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506L7.159 7.451c-.647-.566-1.659-.106-1.659.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z"/><path d="M13.5 17.796L18.981 13L13.5 8.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506l-5.482-4.796c-.646-.566-1.658-.106-1.658.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z"/><path d="M20 7.5a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-1 0V8a.5.5 0 0 1 .5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilNextTrackCircleFilled0)"/></g>`),
		g.Group(children),
	)
}