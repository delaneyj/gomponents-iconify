package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.948 15.145a.5.5 0 0 1-.453-.543l.471-5.186a.5.5 0 0 1 .996.09l-.471 5.186a.5.5 0 0 1-.543.453Z"/><path d="M15.148 8.945a.5.5 0 0 1-.453.543L9.51 9.96a.5.5 0 0 1-.09-.996l5.185-.472a.5.5 0 0 1 .543.453Z"/><path d="M9.647 9.643a.5.5 0 0 1 .707 0l6.535 6.536a.5.5 0 1 1-.707.707l-6.535-6.535a.5.5 0 0 1 0-.708Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}