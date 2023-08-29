package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marvelapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.339 8.13c1.373 0-1.162 7.076-.845 10.138c.317 3.063 3.696 2.218 3.485.423c-.423-3.063 1.69-12.672 3.696-12.672c1.478 0-1.69 6.547-1.056 10.665c.422 2.64 4.012 1.901 3.59.106c-1.162-5.386 2.64-10.56 2.112-14.361C21.11.845 20.159 0 19.209 0c-3.379 0-6.125 6.97-6.125 6.97s.423-3.908-2.428-4.119C6.643 2.64 2.525 12.777 2.63 21.964c.106 2.957 3.696 2.429 3.485.106c-.211-4.12 2.112-13.94 4.225-13.94z"/>`),
		g.Group(children),
	)
}