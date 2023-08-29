package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a3 3 0 0 0-3 3v6a3 3 0 1 0 6 0V6a3 3 0 0 0-3-3Zm0-2a5 5 0 0 1 5 5v6a5 5 0 0 1-10 0V6a5 5 0 0 1 5-5ZM2.19 13.961l1.962-.392a8.003 8.003 0 0 0 15.692 0l1.962.392C20.895 18.545 16.85 22 11.999 22s-8.896-3.455-9.808-8.039Z"/>`),
		g.Group(children),
	)
}