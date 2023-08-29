package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.507 9.04a3 3 0 0 0-2.466 3.453l4 24a3 3 0 0 0 5.564.995L24 28.047l5.395 9.441a3 3 0 0 0 5.564-.995l4-24a3 3 0 1 0-5.918-.986l-2.616 15.691l-3.82-6.687a3 3 0 0 0-5.21 0l-3.82 6.687l-2.616-15.691a3 3 0 0 0-3.452-2.466Zm-.494 3.124a1 1 0 0 1 1.973-.329l3.077 18.461a1 1 0 0 0 1.855.332l5.213-9.124a1 1 0 0 1 1.737 0l5.214 9.124a1 1 0 0 0 1.854-.332l3.077-18.46a1 1 0 0 1 1.973.328l-4 24a1 1 0 0 1-1.854.332l-6.264-10.961a1 1 0 0 0-1.736 0l-6.264 10.961a1 1 0 0 1-1.855-.332l-4-24Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}