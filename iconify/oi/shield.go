package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="m4 0l-.19.09l-3.5 1.47l-.31.13V2c0 1.66.67 3.12 1.47 4.19c.4.53.83.97 1.25 1.28c.42.31.83.53 1.28.53c.46 0 .86-.22 1.28-.53c.42-.31.85-.75 1.25-1.28C7.33 5.12 8 3.66 8 2v-.31l-.31-.13L4.19.09L4 0zm0 1.09V7c-.04 0-.33-.07-.66-.31s-.71-.63-1.06-1.09a6.26 6.26 0 0 1-1.22-3.28L4 1.1z"/>`),
		g.Group(children),
	)
}