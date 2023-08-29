package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006.23 623l-261 264q-2 2-13 8q-10-6-12-8l-16-14l224-226q32-33 32-65q0-30-6-51.5t-11.5-29t-14.5-16.5l-448-453q-10-10-30-21t-29-11h62q28 1 45 18l478 483q18 19 18 61.5t-18 60.5zm-142-9l-256 259q-19 19-48.5 19t-47.5-19l-494-501q-25-26-15-62q-3-10-3-19V65q0-27 19-46t45-19h224q9 0 18 3q37-10 62 15l496 499q18 19 18 49t-18 48zm-640-485q-40 0-68 28.5t-28 69t28 68.5t68 28t68-28t28-68.5t-28-69t-68-28.5z"/>`),
		g.Group(children),
	)
}