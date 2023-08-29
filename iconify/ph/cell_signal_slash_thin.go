package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 152v48a4 4 0 0 1-8 0v-48a4 4 0 0 1 8 0Zm-44 36a4 4 0 0 0-4 4v8a4 4 0 0 0 8 0v-8a4 4 0 0 0-4-4Zm171 25.31l-160-176a4 4 0 0 0-6 5.38l71 78.06V200a4 4 0 0 0 8 0v-70.45l32 35.2V200a4 4 0 0 0 8 0v-26.45l41 45.14a4 4 0 1 0 5.92-5.38Zm-51-93.68a4 4 0 0 0 4-4V72a4 4 0 0 0-8 0v43.63a4 4 0 0 0 4 4Zm40 44a4 4 0 0 0 4-4V32a4 4 0 0 0-8 0v127.63a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}