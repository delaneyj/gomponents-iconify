package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLoopCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFileLoopCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 7a2.5 2.5 0 0 1 2.5-2.5h5.1a1 1 0 0 1 .702.288l4.4 4.333a1 1 0 0 1 .298.712V17a2.5 2.5 0 0 1-2.5 2.5H12a1 1 0 1 1 0-2h6.5a.5.5 0 0 0 .5-.5v-6.167h-2.4a2 2 0 0 1-2-2V6.5h-4.1a.5.5 0 0 0-.5.5v1.5a1 1 0 0 1-2 0V7Zm8.6.888l.96.945h-.96v-.945Z"/><path d="M11.049 11.678a2.193 2.193 0 1 0-2.058 3.873a2.193 2.193 0 0 0 2.058-3.873Zm-4.732-.031a4.193 4.193 0 1 1 2.674 6.033l-1.676 3.155a1 1 0 0 1-1.767-.938l1.677-3.155a4.195 4.195 0 0 1-.908-5.095Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFileLoopCircleFilled0)"/></g>`),
		g.Group(children),
	)
}