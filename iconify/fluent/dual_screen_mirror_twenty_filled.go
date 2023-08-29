package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenMirrorTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 16H16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-5.5v12Zm-1-12H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h5.5V4ZM5.75 8a.5.5 0 0 1 .432.248l1.75 3A.5.5 0 0 1 7.5 12H4a.5.5 0 0 1-.432-.752l1.75-3A.5.5 0 0 1 5.75 8Zm8.932.248l1.75 3A.5.5 0 0 1 16 12h-3.5a.5.5 0 0 1-.432-.752l1.75-3a.5.5 0 0 1 .864 0Z"/>`),
		g.Group(children),
	)
}