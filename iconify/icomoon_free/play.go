package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.331 2.502C13.087 2.179 10.607 2 8 2s-5.087.179-7.331.502C.239 4.185 0 6.045 0 8s.239 3.815.669 5.498C2.913 13.821 5.393 14 8 14s5.087-.179 7.331-.502C15.761 11.815 16 9.955 16 8s-.239-3.815-.669-5.498zM6 11V5l5 3l-5 3z"/>`),
		g.Group(children),
	)
}