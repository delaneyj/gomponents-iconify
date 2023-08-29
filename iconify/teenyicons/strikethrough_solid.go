package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.781A2.781 2.781 0 0 1 6.781 1H8a3 3 0 0 1 3 3h-1a2 2 0 0 0-2-2H6.781A1.78 1.78 0 0 0 5 3.781c0 .843.502 1.605 1.277 1.937l-.394.919A3.107 3.107 0 0 1 4 3.78ZM9.392 8H2V7h11v1h-2.01c.123.14.237.287.34.44c.405.602.67 1.326.67 2.047A3.513 3.513 0 0 1 8.487 14H7a4 4 0 0 1-4-4h1a3 3 0 0 0 3 3h1.487A2.513 2.513 0 0 0 11 10.487c0-.484-.182-1.017-.5-1.488c-.296-.44-.69-.797-1.108-.999Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}