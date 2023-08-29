package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KickstarterK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 5a3 3 0 0 0-3 3v16a3 3 0 1 0 6 0v-4.586l5.742 6.563A2.996 2.996 0 0 0 22.002 27a2.999 2.999 0 0 0 2.256-4.975l-5.399-6.17L23.4 9.801a3.001 3.001 0 0 0-4.8-3.602L14 12.334V8a3 3 0 0 0-3-3z"/>`),
		g.Group(children),
	)
}