package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugBeetleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 142h18a6 6 0 0 0 0-12h-18v-20h18a6 6 0 0 0 0-12h-18.25a77.81 77.81 0 0 0-22.38-48.88l20.87-20.88a6 6 0 0 0-8.48-8.48l-21.49 21.49a77.8 77.8 0 0 0-92.53 0l-21.5-21.49a6 6 0 0 0-8.48 8.48l20.88 20.88A77.76 77.76 0 0 0 50.25 98H32a6 6 0 0 0 0 12h18v20H32a6 6 0 0 0 0 12h18v10a78.6 78.6 0 0 0 .66 10H32a6 6 0 0 0 0 12h21.18a78 78 0 0 0 149.64 0H224a6 6 0 0 0 0-12h-18.66a78.6 78.6 0 0 0 .66-10ZM128 38a66.07 66.07 0 0 1 65.71 60H62.29A66.07 66.07 0 0 1 128 38Zm6 179.71V136a6 6 0 0 0-12 0v81.71A66.07 66.07 0 0 1 62 152v-42h132v42a66.07 66.07 0 0 1-60 65.71Z"/>`),
		g.Group(children),
	)
}