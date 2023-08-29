package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCompass0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m12 44l10.688-28.5M36 44L25.312 15.5"/><circle cx="24" cy="12" r="4" fill="#fff"/><path d="M37.57 33A24.89 24.89 0 0 1 24 37c-5.003 0-9.662-1.47-13.57-4M24 8V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCompass0)"/>`),
		g.Group(children),
	)
}