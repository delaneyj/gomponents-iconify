package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M42.7 486.7h42.7v-256H42.7v256zM469.3 17.3h-42.7v256h42.7v-256zm-384 0H42.7V60h42.7V17.3zm192 0h-42.7v149.3h42.7V17.3zM0 188h128v-85.3H0V188zm21.3-64h85.3v42.7H21.3V124zm213.4 362.7h42.7V337.3h-42.7v149.4zm192 0h42.7V444h-42.7v42.7zM384 316v85.3h128V316H384zm106.7 64h-85.3v-42.7h85.3V380zM192 294.7h128v-85.3H192v85.3zm21.3-64h85.3v42.7h-85.3v-42.7z"/>`),
		g.Group(children),
	)
}