package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0zm5.66 14.24l-1.41 1.41L10 11.41l-4.24 4.25l-1.42-1.42L8.59 10L4.34 5.76l1.42-1.42L10 8.59l4.24-4.24l1.41 1.41L11.41 10z"/>`),
		g.Group(children),
	)
}