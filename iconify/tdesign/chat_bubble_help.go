package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatBubbleHelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.959 10.959 0 0 1 1 12Zm11-4.5a2 2 0 0 0-1.886 1.333l-.334.943l-1.885-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25h-2V14.5c0-.867.39-1.573.826-2.11c.432-.53.974-.974 1.41-1.318A2 2 0 0 0 12 7.5Zm-1 9.25h2.004v2.004H11V16.75Z"/>`),
		g.Group(children),
	)
}