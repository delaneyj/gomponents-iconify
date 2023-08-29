package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOutlineAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q21-40 59-63.5T256 21q58 0 102 37t55 92zm-8 170q27 0 45.5-19t18.5-45t-18.5-45t-45.5-19h-32v-11q0-48-34.5-82.5T256 64q-58 0-94 47q41 12 67.5 46t26.5 78h-43q0-36-25-61t-60-25t-60 25t-25 60.5T68 295t60 25h277z"/>`),
		g.Group(children),
	)
}