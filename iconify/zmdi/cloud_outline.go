package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q20-40 58-63.5T256 21q58 0 102 37t55 92zm-8 170q27 0 45.5-18.5T469 256t-18.5-45.5T405 192h-32v-11q0-48-34-82.5T256 64q-40 0-71 24t-42 61h-15q-35 0-60 25t-25 60.5T68 295t60 25h277z"/>`),
		g.Group(children),
	)
}