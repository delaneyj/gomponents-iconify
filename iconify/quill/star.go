package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.546 4.984a.5.5 0 0 1 .908 0l3.097 6.714a.5.5 0 0 0 .395.287l7.341.87a.5.5 0 0 1 .28.864l-5.427 5.02a.5.5 0 0 0-.15.464l1.44 7.251a.5.5 0 0 1-.735.534l-6.45-3.611a.5.5 0 0 0-.49 0l-6.45 3.61a.5.5 0 0 1-.735-.533l1.44-7.251a.5.5 0 0 0-.15-.465l-5.428-5.02a.5.5 0 0 1 .28-.863l7.342-.87a.5.5 0 0 0 .396-.287l3.096-6.714Z"/>`),
		g.Group(children),
	)
}