package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGreenland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M54 54H10C3.373 54 0 49.075 0 43V32h64v11c0 6.075-3.373 11-10 11"/><path fill="#e6e7e8" d="M0 21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v11H0V21zm18.18 23.715C25.2 44.715 30.896 40 30.896 32H5.464c0 8 5.693 12.715 12.716 12.715"/><path fill="#ec1c24" d="M18.18 19.284C11.16 19.284 5.464 25 5.464 32h25.432c0-7-5.694-12.716-12.716-12.716"/>`),
		g.Group(children),
	)
}