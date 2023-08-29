package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PennantTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14 2a1 1 0 0 1 .993.883L15 3v17h1a1 1 0 0 1 .117 1.993L16 22h-4a1 1 0 0 1-.117-1.993L12 20h1v-7.351L4.594 8.914c-.752-.335-.79-1.365-.113-1.77l.113-.058L13 3.35V3a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}