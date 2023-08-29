package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 448q0 26-19 45L557 941q-19 19-45 19t-45-19L19 493Q0 474 0 448t19-45t45-19h896q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}