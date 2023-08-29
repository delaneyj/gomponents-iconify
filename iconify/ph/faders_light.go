package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 120v96a6 6 0 0 1-12 0v-96a6 6 0 0 1 12 0Zm66 74a6 6 0 0 0-6 6v16a6 6 0 0 0 12 0v-16a6 6 0 0 0-6-6Zm24-32h-18V40a6 6 0 0 0-12 0v122h-18a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm-168 0a6 6 0 0 0-6 6v48a6 6 0 0 0 12 0v-48a6 6 0 0 0-6-6Zm24-32H62V40a6 6 0 0 0-12 0v90H32a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Zm72-48h-18V40a6 6 0 0 0-12 0v42h-18a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}