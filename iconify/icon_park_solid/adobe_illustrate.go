package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeIllustrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAdobeIllustrate0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m20 15l-6 18m18 0v-8m0-5v-1m-12-4l6 18m-10-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAdobeIllustrate0)"/>`),
		g.Group(children),
	)
}