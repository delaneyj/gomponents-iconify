package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405.3 192L320 0h-42.7l85.3 192h42.7zM234.7 0H192l-85.3 192h42.7L234.7 0zm-192 469.3c0 23.5 19.1 42.7 42.7 42.7h341.3c23.5 0 42.7-19.1 42.7-42.7l21.3-192H21.3l21.4 192zm320-149.3h42.7L384 469.3h-42.7L362.7 320zm-128 0h42.7v149.3h-42.7V320zm-85.4 0l21.3 149.3H128L106.7 320h42.6zm341.4-106.7H21.3C9.5 213.3 0 222.9 0 234.7V256h512v-21.3c0-11.8-9.5-21.4-21.3-21.4z"/>`),
		g.Group(children),
	)
}