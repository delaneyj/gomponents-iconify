package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelChair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(2 1)"><ellipse cx="6.49" cy="1.471" rx="1.49" ry="1.471"/><path d="M12.6 14.951a.98.98 0 0 1-.819-.442l-1.272-2.511H5.993a.993.993 0 0 1-.987-.999V5.064c0-.553.442-.998.987-.998c.543 0 .986.445.986.998v4.937h4.056c.329 0 .636.165.817.442l1.084 2.225l.654-.268a.98.98 0 0 1 1.323.449a1.004 1.004 0 0 1-.444 1.338l-1.431.659a.957.957 0 0 1-.438.105z"/><path d="M5.381 15.974c-2.936 0-5.324-2.384-5.324-5.313c0-1.726.845-3.35 2.26-4.346a.667.667 0 0 1 .931.162a.666.666 0 0 1-.161.93a3.983 3.983 0 0 0-1.692 3.254a3.986 3.986 0 0 0 3.987 3.977c1.1 0 2.123-.437 2.885-1.23a.668.668 0 1 1 .964.925a5.292 5.292 0 0 1-3.85 1.641z"/></g>`),
		g.Group(children),
	)
}