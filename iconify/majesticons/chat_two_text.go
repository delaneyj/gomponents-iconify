package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTwoText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 11c0-5.167 5.145-9 11-9s11 3.833 11 9s-5.145 9-11 9c-1.198 0-2.354-.156-3.437-.447c-.785.662-1.59 1.244-2.54 1.672C4.894 21.735 3.617 22 2 22a1 1 0 0 1-.707-1.707c.876-.876 1.843-1.914 2.368-3.416C2.029 15.327 1 13.28 1 11zm7-5a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}