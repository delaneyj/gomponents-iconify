package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddPic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAddPic0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4"><path d="M38 21v19a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V12a2 2 0 0 1 2-2h18.364"/><path fill="#555" d="M12 31.03L18 23l3 3l3.5-5.5L32 31.03H12Z"/><path d="M34 10h8m-4.005-4.205v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAddPic0)"/>`),
		g.Group(children),
	)
}