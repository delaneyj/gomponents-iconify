package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSunset0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="m19 8l14 26H5L19 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m29 26l5-6l9 14H32m-22 7h28"/><circle cx="38" cy="10" r="3" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSunset0)"/>`),
		g.Group(children),
	)
}