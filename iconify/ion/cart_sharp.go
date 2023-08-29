package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="176" cy="416" r="32" fill="currentColor"/><circle cx="400" cy="416" r="32" fill="currentColor"/><path fill="currentColor" d="M167.78 304h261.34l38.4-192H133.89l-8.47-48H32v32h66.58l48 272H432v-32H173.42l-5.64-32z"/>`),
		g.Group(children),
	)
}