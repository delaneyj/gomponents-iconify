package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostSmileLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a9 9 0 0 1 9 9v7.5a3.5 3.5 0 0 1-6.39 1.976a2.999 2.999 0 0 1-5.223 0a3.5 3.5 0 0 1-6.382-1.783L3 18.499V11a9 9 0 0 1 9-9Zm0 2a7 7 0 0 0-6.996 6.76L5 11v7.446l.002.138a1.5 1.5 0 0 0 2.645.88l.088-.116a2 2 0 0 1 3.393.142a.999.999 0 0 0 1.74.003a2 2 0 0 1 3.296-.278l.097.13a1.5 1.5 0 0 0 2.732-.701L19 18.5V11a7 7 0 0 0-7-7Zm4 9a4 4 0 0 1-7.995.2L8 13h2a2 2 0 1 0 4 0h2Zm-4-6a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}