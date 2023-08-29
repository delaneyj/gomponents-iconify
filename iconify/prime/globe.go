package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 9 9a9 9 0 0 0-9-9Zm7.46 8.25H16.7a13 13 0 0 0-2.94-6.53a7.52 7.52 0 0 1 5.7 6.53Zm-10.65 1.5h6.38A13.18 13.18 0 0 1 12 19.1a13.18 13.18 0 0 1-3.19-6.35Zm0-1.5A13.18 13.18 0 0 1 12 4.9a13.18 13.18 0 0 1 3.19 6.35Zm1.43-6.53a13 13 0 0 0-2.94 6.53H4.54a7.52 7.52 0 0 1 5.7-6.53Zm-5.7 8H7.3a13 13 0 0 0 2.94 6.53a7.52 7.52 0 0 1-5.7-6.5Zm9.22 6.53a13 13 0 0 0 2.94-6.53h2.76a7.52 7.52 0 0 1-5.7 6.56Z"/>`),
		g.Group(children),
	)
}