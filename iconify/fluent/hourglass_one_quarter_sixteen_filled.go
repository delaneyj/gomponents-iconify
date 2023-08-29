package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassOneQuarterSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.113 11a8.75 8.75 0 0 0-.095.08c-.296.257-.518.496-.683.806c-.13.247-.256.595-.31 1.114h5.946c-.054-.52-.179-.868-.31-1.114c-.164-.31-.386-.55-.681-.806a7.517 7.517 0 0 0-.095-.08H6.113Z"/>`),
		g.Group(children),
	)
}