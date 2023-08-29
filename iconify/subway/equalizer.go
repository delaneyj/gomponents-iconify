package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M106.7 25.8H21.3v64h85.3v-64zM21.3 495.2h85.3V217.8H21.3v277.4zM490.7 25.8h-85.3v277.3h85.3V25.8zm-85.4 469.4h85.3v-64h-85.3v64zm-192 0h85.3V324.5h-85.3v170.7zm85.4-469.4h-85.3v170.7h85.3V25.8zM0 175.2h128v-42.7H0v42.7zm192 106.6h128v-42.7H192v42.7zm192 64v42.7h128v-42.7H384z"/>`),
		g.Group(children),
	)
}