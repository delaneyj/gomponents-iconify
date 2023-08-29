package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialGoogle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M707 360c-70-75-128-87-184-88h-2c-126 0-245 104-245 248c0 151 136 239 244 239h1c60 0 133-11 197-103H472V401l523 2c5 26 14 91 14 125c0 289-194 495-493 495C232 1023 1 797 1 515S232 8 516 8c139 0 288 55 382 180zm-187 93v154h263c-12 65-81 195-263 195c-159 0-287-130-287-285c0-156 131-287 287-287c91 0 152 40 185 72l126-119C751 108 646 63 520 63C264 63 56 264 56 517c0 251 208 453 464 453c270 0 445-185 445-442c0-29-2-52-6-75H520z"/>`),
		g.Group(children),
	)
}