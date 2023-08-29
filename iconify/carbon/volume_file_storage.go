package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeFileStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25.707 17.293l-5-5A1 1 0 0 0 20 12h-6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V18a1 1 0 0 0-.293-.707ZM23.586 18H20v-3.586ZM14 28V14h4v4a2 2 0 0 0 2 2h4v8Z"/><path fill="currentColor" d="M8 27H4a2.002 2.002 0 0 1-2-2V5a2.002 2.002 0 0 1 2-2h7.586A1.986 1.986 0 0 1 13 3.586L16.414 7H28a2.002 2.002 0 0 1 2 2v8h-2V9H15.586l-4-4H4v20h4Z"/>`),
		g.Group(children),
	)
}