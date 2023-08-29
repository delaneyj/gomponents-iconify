package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aquarium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.5 19.5c-4.97 0-9-4.253-9-9.5c0-1.294.245-2.527.689-3.65c0 0-.995-.85-3.189-.85L.75 5S2.5 1.5 7.705 1.959C9.093 1.035 11.238.5 13 .5c3.61 0 5.263 1.353 6.571 3.384c.664 1.03 1.446 1.989 2.441 2.706l.488.351V8.5H14c0 3-.877 5.393-2.5 6.33l-.5-.33s.5-3-.92-3.934c-.968.866-1.58 2.128-1.58 3.934c0 2.761 1.79 5 4 5Zm0 0c1.198-2.207 2.5-4 5-4v.286c-.74 1.025-1.5 2.37-1.5 3.714c0 1.344.76 2.689 1.5 3.714v.286c-2.5 0-3.802-1.793-5-4Zm3.5-16v2"/>`),
		g.Group(children),
	)
}