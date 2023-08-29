package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm-7 8a7 7 0 1 1 14 0a7 7 0 0 1-14 0Zm5.303-3.554A.5.5 0 0 1 8.8 6H12a.5.5 0 1 1 0 1H9.248l-.19 1.74l.246-.002c.42-.003.95.002 1.221.04a2.875 2.875 0 1 1-2.972 4.132a.5.5 0 0 1 .894-.447a1.875 1.875 0 1 0 1.939-2.694c-.184-.027-.632-.034-1.076-.031a40.593 40.593 0 0 0-.733.01l-.047.002h-.016a.5.5 0 0 1-.511-.554l.3-2.75Z"/>`),
		g.Group(children),
	)
}