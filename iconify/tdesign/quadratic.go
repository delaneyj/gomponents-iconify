package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quadratic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.253 3H23v2H12.747L7.882 21.418l-4.215-6.775H1v-2h3.778l2.451 3.94L11.253 3ZM22 10v2.136a2 2 0 0 1-.726 1.542L19.07 15.5l2.204 1.822A2 2 0 0 1 22 18.864V21h-2v-2.136l-2.5-2.067l-2.5 2.067V21h-2v-2.136a2 2 0 0 1 .726-1.542L15.93 15.5l-2.204-1.822A2 2 0 0 1 13 12.136V10h2v2.136l2.5 2.067l2.5-2.067V10h2Z"/>`),
		g.Group(children),
	)
}