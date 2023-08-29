package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPaperPlaneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="m3.874 10.454l7.823 2.046l2.803 7.868a.5.5 0 0 0 .95-.026l4.745-16.085a.5.5 0 0 0-.654-.61L3.826 9.502a.5.5 0 0 0 .048.952Zm1.783-.567l13.296-4.954l-4.027 13.652l-2.376-6.67a.5.5 0 0 0-.344-.315L5.657 9.887Z" clip-rule="evenodd"/><path d="m19 4.293l.707.707L12 12.707L11.293 12L19 4.293Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPaperPlaneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}