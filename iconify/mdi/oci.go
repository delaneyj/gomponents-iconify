package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oci(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v20h20V2m-2.88 17.03H4.87V5h14.26v14.03m-4.3-8.32h2.86v6.88h-2.86m0-11.18h2.86v2.86h-2.86M6.3 6.41v11.18h7.1v-2.87H9.17V9.28h4.23V6.41Z"/>`),
		g.Group(children),
	)
}