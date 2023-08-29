package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm6.8-4a.5.5 0 0 0-.497.446l-.3 2.75a.5.5 0 0 0 .51.554h.017l.047-.002a37.818 37.818 0 0 1 .733-.01c.444-.003.892.004 1.076.03a1.875 1.875 0 1 1-1.939 2.694a.5.5 0 1 0-.894.448a2.875 2.875 0 1 0 2.972-4.132c-.271-.038-.8-.043-1.22-.04c-.086 0-.168 0-.246.002L9.249 7H12a.5.5 0 1 0 0-1H8.8Z"/>`),
		g.Group(children),
	)
}