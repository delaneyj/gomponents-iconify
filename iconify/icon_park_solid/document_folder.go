package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDocumentFolder0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M32 6H22v36h10V6Zm10 0H32v36h10V6ZM10 6l8 1l-3.5 35L6 41l4-35Z"/><path stroke="#000" stroke-linecap="round" d="M37 18v-3m-10 3v-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDocumentFolder0)"/>`),
		g.Group(children),
	)
}