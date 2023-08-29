package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 128H448v832q0 27-19 45.5t-45 18.5t-45-18.5t-19-45.5V128H64q-27 0-45.5-18.5T0 64.5T18.5 19T64 0h640q26 0 45 19t19 45.5t-18.5 45T704 128z"/>`),
		g.Group(children),
	)
}