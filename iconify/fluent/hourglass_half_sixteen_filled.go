package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassHalfSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.025 13h5.946c-.054-.52-.179-.868-.31-1.114c-.164-.31-.386-.55-.681-.806c-.11-.095-.209-.174-.323-.267a37.456 37.456 0 0 1-.186-.151a6.281 6.281 0 0 1-.613-.56C8.401 9.614 8 8.937 8 8c0 .938-.402 1.614-.86 2.103a6.283 6.283 0 0 1-.613.56l-.186.15c-.114.093-.213.173-.323.268c-.296.256-.518.495-.683.805c-.13.247-.256.595-.31 1.114Z"/>`),
		g.Group(children),
	)
}