package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForChile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M0 32v11c0 6.075 3.373 11 10 11h44c6.627 0 10-4.925 10-11V32H0z"/><path fill="#e6e7e8" d="M54 10H23v22h41V21c0-6.075-3.373-11-10-11"/><path fill="#2b3990" d="M23 10H10C3.373 10 0 14.925 0 21v11h23V10z"/><path fill="#e6e7e8" d="m11.52 13.997l2.247 4.554l5.03.739l-3.638 3.548l.86 5l-4.5-2.364l-4.5 2.364l.86-5L4.24 19.29l5.03-.739z"/>`),
		g.Group(children),
	)
}