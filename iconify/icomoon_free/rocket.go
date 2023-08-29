package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 1L6 6H3l-3 4s3.178-.885 5.032-.47L0 16l6.592-5.127C7.511 12.977 6 16 6 16l4-3v-3l5-5l1-5l-5 1z"/>`),
		g.Group(children),
	)
}