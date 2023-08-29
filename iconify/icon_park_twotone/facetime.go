package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facetime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFacetime0"><g fill="#555" stroke="#fff" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linejoin="round" d="M24 19H12v10h12V19Zm11-1l-5 4v4l5 4V18Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFacetime0)"/>`),
		g.Group(children),
	)
}