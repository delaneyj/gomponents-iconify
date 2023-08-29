package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webhook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 26a3 3 0 1 0-2.816-4H13v1a5 5 0 1 1-5-5v-2a7 7 0 1 0 6.929 8h6.255A2.991 2.991 0 0 0 24 26Z"/><path fill="currentColor" d="M24 16a7.024 7.024 0 0 0-2.57.487l-3.166-5.54a3.047 3.047 0 1 0-1.732 1l4.119 7.208l.868-.498a5 5 0 1 1-1.85 6.842l-1.732 1.002A7 7 0 1 0 24 16Z"/><path fill="currentColor" d="M8.532 20.054a3.03 3.03 0 1 0 1.733.998l3.625-6.344l.498-.868l-.868-.497a5 5 0 1 1 6.812-1.844l1.731 1.002a7 7 0 1 0-10.346 2.036c-.457.742-1.102 1.871-2.073 3.572Z"/>`),
		g.Group(children),
	)
}