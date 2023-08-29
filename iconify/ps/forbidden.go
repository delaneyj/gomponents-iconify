package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forbidden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zM43 256q0-88 62.5-150.5T256 43q79 0 134 49L92 390q-49-55-49-134zm213 213q-79 0-134-49l298-298q49 55 49 134q0 88-62.5 150.5T256 469z"/>`),
		g.Group(children),
	)
}