package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightToBracket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m217.9 105.9l122.8 122.8c7.2 7.2 11.3 17.1 11.3 27.3s-4.1 20.1-11.3 27.3L217.9 406.1c-6.4 6.4-15 9.9-24 9.9c-18.7 0-33.9-15.2-33.9-33.9V320H32c-17.7 0-32-14.3-32-32v-64c0-17.7 14.3-32 32-32h128v-62.1c0-18.7 15.2-33.9 33.9-33.9c9 0 17.6 3.6 24 9.9zM352 416h64c17.7 0 32-14.3 32-32V128c0-17.7-14.3-32-32-32h-64c-17.7 0-32-14.3-32-32s14.3-32 32-32h64c53 0 96 43 96 96v256c0 53-43 96-96 96h-64c-17.7 0-32-14.3-32-32s14.3-32 32-32z"/>`),
		g.Group(children),
	)
}