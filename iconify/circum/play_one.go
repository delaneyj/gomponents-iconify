package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.562 21.94a2.5 2.5 0 0 1-2.5-2.5V4.56A2.5 2.5 0 0 1 7.978 2.5l10.877 7.439a2.5 2.5 0 0 1 0 4.12L7.977 21.5a2.5 2.5 0 0 1-1.415.44Zm0-18.884a1.494 1.494 0 0 0-.7.177a1.477 1.477 0 0 0-.8 1.327v14.879a1.5 1.5 0 0 0 2.35 1.235l10.877-7.44a1.5 1.5 0 0 0 0-2.471L7.413 3.326a1.491 1.491 0 0 0-.849-.27Z"/>`),
		g.Group(children),
	)
}