package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.25 9v5a2 2 0 0 0 2 2h7.5a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-7.5a2 2 0 0 0-2 2Zm-1 0v5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Z"/><path d="M11.5 11.5L9 13v-3l2.5 1.5Z"/><path d="M12 11.5a.5.5 0 0 1-.243.429l-2.5 1.5A.5.5 0 0 1 8.5 13v-3a.5.5 0 0 1 .757-.429l2.5 1.5A.5.5 0 0 1 12 11.5Zm-2.5-.617v1.234l1.028-.617l-1.028-.617Zm3.062-7.773a.5.5 0 0 1 .078.702l-2 2.5a.5.5 0 1 1-.78-.624l2-2.5a.5.5 0 0 1 .702-.078Z"/><path d="M6.938 3.11a.5.5 0 0 1 .702.078l2 2.5a.5.5 0 1 1-.78.624l-2-2.5a.5.5 0 0 1 .078-.702Z"/></g>`),
		g.Group(children),
	)
}