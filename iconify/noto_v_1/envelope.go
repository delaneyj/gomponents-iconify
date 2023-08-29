package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#fcc21b" fill-rule="evenodd" clip-rule="evenodd"><path d="M20.97 50.92v45.02h86.06V50.92L64 72.31z"/><path d="M20.97 32.06v11.98L64 65.43l43.03-21.39V32.06z"/></g>`),
		g.Group(children),
	)
}