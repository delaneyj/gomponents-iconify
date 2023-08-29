package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMask0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M14 14h20c5.523 0 10 4.477 10 10v3c0 7.18-5.82 13-13 13H17C9.82 40 4 34.18 4 27v-3c0-5.523 4.477-10 10-10Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M18 27h12m-6-6v12"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M4 25.042V12.014A4.168 4.168 0 0 1 11.047 9c1.6 1.527 2.74 3.194 3.424 5M44 25.042V12.766A4.478 4.478 0 0 0 36 10a111.149 111.149 0 0 0-3 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMask0)"/>`),
		g.Group(children),
	)
}