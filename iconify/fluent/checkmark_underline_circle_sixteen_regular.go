package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkUnderlineCircleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 13A5 5 0 1 1 8 3a5 5 0 0 1 0 10Zm0 1A6 6 0 1 0 8 2a6 6 0 0 0 0 12Zm2.856-8.85a.521.521 0 0 1 0 .719L7.972 8.85a.484.484 0 0 1-.696 0L6.144 7.68a.521.521 0 0 1 0-.72a.48.48 0 0 1 .696 0l.784.81l2.536-2.62a.48.48 0 0 1 .696 0ZM6 10.5a.5.5 0 0 1 .5-.5h2.998a.5.5 0 0 1 0 1H6.5a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}