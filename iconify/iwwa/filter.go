package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M32.008 5.426a.121.121 0 0 0-.131-.032L7.749 13.566a.125.125 0 0 0-.017.229l12.992 6.685v14.128c0 .047.026.09.069.111a.125.125 0 0 0 .056.014a.126.126 0 0 0 .074-.024l5.338-3.932a.125.125 0 0 0 .051-.101V19.67l5.721-14.111a.128.128 0 0 0-.025-.133zM25.37 19.391a.123.123 0 0 0-.009.047l-.027 10.809l-3.632 2.676V19.961a.124.124 0 0 0-.068-.111L9.968 13.847l20.459-6.93l-5.057 12.474z"/>`),
		g.Group(children),
	)
}