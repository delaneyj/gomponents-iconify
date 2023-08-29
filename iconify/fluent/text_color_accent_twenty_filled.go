package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColorAccentTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 13.125c0-.345.243-.625.542-.625h11.916c.3 0 .542.28.542.625v3.75c0 .345-.242.625-.542.625H4.042c-.3 0-.542-.28-.542-.625v-3.75Z"/>`),
		g.Group(children),
	)
}