package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.089.064H1.913A.912.912 0 0 0 1 .975V11.05c0 .504.409.912.913.912h14.176A.911.911 0 0 0 17 11.05V.975a.912.912 0 0 0-.911-.911zM2.971 9.369V2.575c0-.375.283-.677.633-.677h10.794c.351 0 .633.302.633.677v6.794c0 .373-.282.676-.633.676H3.604c-.35 0-.633-.303-.633-.676zm11.476 4.714H9.911v-2.041H8.052v2.041H3.583c-.282 0-.511.792-.511 1.193c0 .403.229.623.511.623h10.864c.282 0 .511-.22.511-.623c0-.401-.229-1.193-.511-1.193z"/>`),
		g.Group(children),
	)
}