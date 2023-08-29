package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipUpload0" width="48" height="48" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:alpha"><path fill="#fff" d="M48 0H0V48H48V0Z"/></mask><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" mask="url(#ipUpload0)"><path d="M6 24.0083V42H42V24"/><path d="M33 15L24 6L15 15"/><path d="M23.9917 32V6"/></g>`),
		g.Group(children),
	)
}