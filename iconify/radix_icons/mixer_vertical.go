package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MixerVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 1.5a.5.5 0 0 0-1 0V7c0 .017 0 .033.002.05a2.5 2.5 0 0 0 0 4.9A.506.506 0 0 0 4 12v1.5a.5.5 0 0 0 1 0V12c0-.017 0-.033-.002-.05a2.5 2.5 0 0 0 0-4.9A.507.507 0 0 0 5 7V1.5Zm6 0a.5.5 0 0 0-1 0V3c0 .017 0 .033.002.05a2.5 2.5 0 0 0 0 4.9A.507.507 0 0 0 10 8v5.5a.5.5 0 0 0 1 0V8c0-.017 0-.033-.002-.05a2.5 2.5 0 0 0 0-4.9A.507.507 0 0 0 11 3V1.5ZM4.5 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM9 5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}