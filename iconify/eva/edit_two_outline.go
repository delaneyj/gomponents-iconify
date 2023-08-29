package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaEdit2Outline0"><g id="evaEdit2Outline1"><path id="evaEdit2Outline2" fill="currentColor" d="M19 20H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2ZM5 18h.09l4.17-.38a2 2 0 0 0 1.21-.57l9-9a1.92 1.92 0 0 0-.07-2.71L16.66 2.6A2 2 0 0 0 14 2.53l-9 9a2 2 0 0 0-.57 1.21L4 16.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 18ZM15.27 4L18 6.73l-2 1.95L13.32 6Zm-8.9 8.91L12 7.32l2.7 2.7l-5.6 5.6l-3 .28Z"/></g></g>`),
		g.Group(children),
	)
}