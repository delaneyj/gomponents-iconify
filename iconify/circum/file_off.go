package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3.308a.5.5 0 0 0-.7.71l.76.76v14.67a2.5 2.5 0 0 0 2.5 2.5h10.88a2.476 2.476 0 0 0 2.28-1.51l.28.28c.45.45 1.16-.26.7-.71Zm14.92 16.33a1.492 1.492 0 0 1-1.48 1.31H6.56a1.5 1.5 0 0 1-1.5-1.5V5.778Zm-5.54-16.55v2.92a2.5 2.5 0 0 0 2.5 2.5h3.07l-.01 6.7a.5.5 0 0 0 1 0v-6.67a2.057 2.057 0 0 0-.75-1.47c-1.3-1.26-2.59-2.53-3.89-3.8a3.924 3.924 0 0 0-1.41-1.13a6.523 6.523 0 0 0-1.71-.06H6.81a.5.5 0 0 0 0 1Zm4.83 4.42h-2.33a1.5 1.5 0 0 1-1.5-1.5v-2.24Z"/>`),
		g.Group(children),
	)
}