package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmbassyEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 2a4.07 4.07 0 0 0-2.41 1.06a.36.36 0 0 0-.09.24v3.32a.34.34 0 0 0 .57.27a3.23 3.23 0 0 1 1.93-.67C6.61 6.22 6.85 7 8 7a3.08 3.08 0 0 0 1.83-.84a.36.36 0 0 0 .17-.32V2.37A.35.35 0 0 0 9.5 2a3.13 3.13 0 0 1-1.5.67C6.85 2.67 6.65 2 5.5 2zm-4-.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm.5 3v6a.5.5 0 0 1-1 0v-6a.5.5 0 0 1 1 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}