package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumanHandsdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1a2 2 0 0 0-2 2c0 1.11.89 2 2 2c1.11 0 2-.89 2-2a2 2 0 0 0-2-2m-2 5c-.27 0-.5.11-.69.28H9.3L4 11.59L5.42 13L9 9.41V22h2v-7h2v7h2V9.41L18.58 13L20 11.59l-5.3-5.31c-.2-.17-.43-.28-.7-.28"/>`),
		g.Group(children),
	)
}