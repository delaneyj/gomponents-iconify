package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUpload0" width="48" height="48" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:alpha"><path d="M48 0H0v48h48V0Z"/></mask><mask id="ipSUpload1"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" mask="url(#ipSUpload0)"><path d="M6 24.008V42h36V24m-9-9l-9-9l-9 9m8.992 17V6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUpload1)"/>`),
		g.Group(children),
	)
}