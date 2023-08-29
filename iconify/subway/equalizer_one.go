package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M85.3 21.3H42.7V64h42.7V21.3zM42.7 490.7h42.7v-256H42.7v256zm384 0h42.7V448h-42.7v42.7zm42.6-469.4h-42.7v256h42.7v-256zm-192 0h-42.7v149.3h42.7V21.3zm-42.6 469.4h42.7V341.3h-42.7v149.4zM0 192h128v-85.3H0V192zm192 106.7h128v-85.3H192v85.3zM384 320v85.3h128V320H384z"/>`),
		g.Group(children),
	)
}