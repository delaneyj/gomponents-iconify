package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabasePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.43 5C10.981 5 14 3.837 14 2.5S10.981 0 7.43 0S1 1.163 1 2.5S3.879 5 7.43 5z"/><path d="M7.494 9.919c3.561 0 6.447-.982 6.447-2.196V4.377c0 .672-2.932 1.674-6.447 1.674S1.047 5.049 1.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196zM11 13h4.915v.957H11z"/><path d="M1.016 9.444v3.269c0 1.188 2.887 2.146 6.447 2.146c.869 0 1.697-.059 2.455-.162v-2.739h2.04V10.52c-1.169.321-2.76.561-4.495.561c-3.516-.001-6.447-.981-6.447-1.637zM13 11h.958v4.937H13z"/></g>`),
		g.Group(children),
	)
}