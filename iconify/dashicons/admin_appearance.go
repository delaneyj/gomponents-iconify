package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminAppearance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.48 11.06L7.41 3.99l1.5-1.5c.5-.56 2.3-.47 3.51.32c1.21.8 1.43 1.28 2.91 2.1c1.18.64 2.45 1.26 4.45.85zm-.71.71L6.7 4.7L4.93 6.47a.996.996 0 0 0 0 1.41l1.06 1.06c.39.39.39 1.03 0 1.42c-.6.6-1.43 1.11-2.21 1.69c-.35.26-.7.53-1.01.84C1.43 14.23.4 16.08 1.4 17.07c.99 1 2.84-.03 4.18-1.36c.31-.31.58-.66.85-1.02c.57-.78 1.08-1.61 1.69-2.21a.996.996 0 0 1 1.41 0l1.06 1.06c.39.39 1.02.39 1.41 0z"/>`),
		g.Group(children),
	)
}