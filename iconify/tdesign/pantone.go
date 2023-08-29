package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pantone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.347.991l3.178 4.455l2.099-.715l1.915 5.602H23V22H1V10.333h1.652L14.347.991ZM3.353 12.333l-.003.003H3V20h18v-7.667H3.353Zm16.072-2L18.388 7.3l-1.604.513l-7.397 2.521h10.038ZM15.549 6.12l-1.604-2.248l-6.102 4.874L15.55 6.12ZM5.285 15.164H7.29v2.004H5.285v-2.004Z"/>`),
		g.Group(children),
	)
}