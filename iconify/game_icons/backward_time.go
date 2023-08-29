package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackwardTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m208.242 24.629l-52.058 95.205l95.207 52.059l17.271-31.586l-42.424-23.198A143.26 143.26 0 0 1 256 114c78.638 0 142 63.362 142 142s-63.362 142-142 142s-142-63.362-142-142c0-16.46 2.785-32.247 7.896-46.928l-32.32-16.16C82.106 212.535 78 233.798 78 256c0 98.093 79.907 178 178 178s178-79.907 178-178S354.093 78 256 78c-13.103 0-25.875 1.44-38.18 4.148l22.008-40.25l-31.586-17.27zm104.27 130.379L247 253.275V368h18V258.725l62.488-93.733l-14.976-9.984z"/>`),
		g.Group(children),
	)
}