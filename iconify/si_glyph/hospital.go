package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H3v3.084H2v9.874c0 .547.447.989 1 .989h5V12h3.042v3l4.875-.053c.552 0 1.083-.442 1.083-.989V4h-1V1zM7.031 12.97H4v-1h3.031v1zm0-3H4v-1h3.031v1zm0-3H4v-1h3.031v1zM6.959 4h-3V3h3v1zM11 5.007h-1v.977H9v-.977H8V3.98h1v-.997h1v.997h1v1.026zm.953-2.028h3v1.047h-3V2.979zM15 12.97h-3v-1h3v1zm0-3h-3v-1h3v1zm0-3h-3v-1h3v1z"/>`),
		g.Group(children),
	)
}