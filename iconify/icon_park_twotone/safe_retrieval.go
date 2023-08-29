package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeRetrieval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSafeRetrieval0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M7 9.127L24.008 4L41 9.127v10.515C41 30.695 34.153 40.506 24.003 44C13.849 40.506 7 30.692 7 19.638V9.128Z"/><path fill="#555" d="M24 30a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke-linecap="round" d="m29 29l6 7"/><path d="M41 19.643C41 30.695 34.153 40.507 24.002 44"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSafeRetrieval0)"/>`),
		g.Group(children),
	)
}