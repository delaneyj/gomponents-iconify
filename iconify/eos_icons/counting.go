package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Counting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v20H2z"/><path fill="currentColor" d="M22 4v2H2V4z"/><path fill="currentColor" d="M20 2h2v20h-2z"/><circle cx="8" cy="5" r="2" fill="currentColor"/><path fill="currentColor" d="M22 10v2H2v-2z"/><circle cx="10" cy="11" r="2" fill="currentColor"/><circle cx="16" cy="11" r="2" fill="currentColor"/><path fill="currentColor" d="M22 16v2H2v-2zm2 5v2H0v-2z"/><circle cx="8" cy="17" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}