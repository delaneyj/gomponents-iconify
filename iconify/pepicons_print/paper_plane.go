package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M1.874 9.454a.5.5 0 0 1-.048-.952l15.715-5.855a.5.5 0 0 1 .654.61l-4.744 16.085a.5.5 0 0 1-.95.026L9.696 11.5L1.874 9.454Z" opacity=".2"/><path fill-rule="evenodd" d="M.874 7.454L8.697 9.5l2.803 7.868a.5.5 0 0 0 .95-.026l4.745-16.085a.5.5 0 0 0-.654-.61L.826 6.502a.5.5 0 0 0 .048.952Zm1.783-.567l13.296-4.954l-4.027 13.652l-2.376-6.67a.5.5 0 0 0-.344-.315L2.657 6.887Z" clip-rule="evenodd"/><path d="m16 1.293l.707.707L9 9.707L8.293 9L16 1.293Z"/></g>`),
		g.Group(children),
	)
}