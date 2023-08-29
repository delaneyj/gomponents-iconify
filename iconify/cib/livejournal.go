package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Livejournal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.12 19.593a8.752 8.752 0 0 0-4.505 4.532l5.692 1.177l-1.181-5.709zM3.303 11.088L0 7.803A15.018 15.018 0 0 1 7.765 0h.005l3.292 3.287a14.912 14.912 0 0 1 5.953-1.229c8.276 0 14.984 6.703 14.984 14.969C31.999 25.298 25.291 32 17.015 32C8.734 32 2.026 25.303 2.026 17.027c0-2.12.473-4.105 1.271-5.933l12.187 12.152a15 15 0 0 1 7.767-7.797L11.063 3.288a15.033 15.033 0 0 0-7.765 7.796z"/>`),
		g.Group(children),
	)
}