package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlorinSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M314.7 32c-38.8 0-73.7 23.3-88.6 59.1L170.7 224H64c-17.7 0-32 14.3-32 32s14.3 32 32 32h80L98.9 396.3c-5 11.9-16.6 19.7-29.5 19.7H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h37.3c38.8 0 73.7-23.3 88.6-59.1L213.3 288H320c17.7 0 32-14.3 32-32s-14.3-32-32-32h-80l45.1-108.3c5-11.9 16.6-19.7 29.5-19.7H352c17.7 0 32-14.3 32-32s-14.3-32-32-32h-37.3z"/>`),
		g.Group(children),
	)
}