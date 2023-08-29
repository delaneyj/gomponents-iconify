package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilArrowUpLeftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.948 15.145a.5.5 0 0 1-.453-.543l.471-5.186a.5.5 0 0 1 .996.09l-.471 5.186a.5.5 0 0 1-.543.453Z"/><path d="M15.148 8.945a.5.5 0 0 1-.453.543L9.51 9.96a.5.5 0 0 1-.09-.996l5.185-.472a.5.5 0 0 1 .543.453Z"/><path d="M9.647 9.643a.5.5 0 0 1 .707 0l6.535 6.536a.5.5 0 1 1-.707.707l-6.535-6.535a.5.5 0 0 1 0-.708Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilArrowUpLeftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}