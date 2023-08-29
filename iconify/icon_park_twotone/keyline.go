package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTKeyline0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 6h36v36H6V6Z"/><path fill="#555" d="M36 24c0 6.627-5.373 12-12 12s-12-5.373-12-12s5.373-12 12-12s12 5.373 12 12Z"/><path d="m4 4l40 40m0-40L4 44"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTKeyline0)"/>`),
		g.Group(children),
	)
}