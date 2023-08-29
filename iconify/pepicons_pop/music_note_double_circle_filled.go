package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteDoubleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMusicNoteDoubleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" transform="translate(3 3)"><path fill-rule="evenodd" d="M6.75 3.81a1 1 0 0 1 .906-.996l9-.846a1 1 0 0 1 1.094.996v2.181a1 1 0 0 1-.901.995l-9 .893a1 1 0 0 1-1.099-.995V3.81Z" clip-rule="evenodd"/><ellipse cx="14.75" cy="15" rx="3" ry="2.5"/><ellipse cx="5.75" cy="16" rx="3" ry="2.5"/><path fill-rule="evenodd" d="M15.75 5h2v10h-2V5Zm-9 1h2v10h-2V6Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMusicNoteDoubleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}