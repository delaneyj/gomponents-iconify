package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarkalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024L384 800L64 1024q-27 0-45.5-18.5T0 960V128q0-53 37.5-90.5T128 0h512q53 0 90.5 37.5T768 128v832q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}