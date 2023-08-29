package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSdCard0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M8 12v32h32V4H14l-6 8Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M15 15v6m6-9v6m12-6v6m-6-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSdCard0)"/>`),
		g.Group(children),
	)
}