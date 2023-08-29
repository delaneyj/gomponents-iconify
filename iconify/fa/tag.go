package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M448 320q0-53-37.5-90.5T320 192t-90.5 37.5T192 320t37.5 90.5T320 448t90.5-37.5T448 320zm1067 576q0 53-37 90l-491 492q-39 37-91 37q-53 0-90-37L91 762q-38-37-64.5-101T0 544V128q0-52 38-90t90-38h416q53 0 117 26.5T763 91l715 714q37 39 37 91z"/>`),
		g.Group(children),
	)
}