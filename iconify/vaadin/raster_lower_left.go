package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RasterLowerLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 7h1v1h-1V7zm-2 0h1v1h-1V7zm-2 0h1v1h-1V7zM9 7h1v1H9V7zm5-1h1v1h-1V6zm-2 0h1v1h-1V6zm-2 0h1v1h-1V6zm5-1h1v1h-1V5zm-2 0h1v1h-1V5zm-2 0h1v1h-1V5zm3-1h1v1h-1V4zm-2 0h1v1h-1V4zm3-1h1v1h-1V3zm-2 0h1v1h-1V3zm1-1h1v1h-1V2zm1-1h1v1h-1V1zM7 15h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm5-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm3-1h1v1H6v-1zm-2 0h1v1H4v-1zm3-1h1v1H7v-1zm-2 0h1v1H5v-1zm1-1h1v1H6v-1zm1-1h1v1H7V9zm8 6h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1V9zm-2 0h1v1h-1V9zm-2 0h1v1h-1V9zM9 9h1v1H9V9zm5-1h1v1h-1V8zm-2 0h1v1h-1V8zm-2 0h1v1h-1V8zM8 8h1v1H8V8z"/>`),
		g.Group(children),
	)
}