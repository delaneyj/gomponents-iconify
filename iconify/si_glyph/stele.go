package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stele(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 14v1.99h15.003V14H1zm10-8h.905v.898H11zM5 6h.951v.905H5z"/><path d="M2 5.412v7.589h13V5.412C14.999.068 10.014.024 8.06.024C6.106.024 2 .068 2 5.412zm4.967 4.574H5.894V7.954h-.887v2.003h-1.04V4.986h2.986v2.028H6.02v.86h.947v2.112zm1.986-.032h-1v-5h1v5zm2.085-1.958v1.958H10v-5h2.995L13 7.984l-1.962.012z"/></g>`),
		g.Group(children),
	)
}