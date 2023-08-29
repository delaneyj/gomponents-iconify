package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LensAlignment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLensAlignment0"><g fill="none"><path fill="#fff" d="M14 10a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 10a4 4 0 0 1-4 4m4-4a4 4 0 1 0-4 4m4-4h6m-10 4v6"/><path fill="#fff" d="M14 38a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 38a4 4 0 0 0-4-4m4 4a4 4 0 1 1-4-4m4 4h6m-10-4v-6"/><path fill="#fff" d="M34 38a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 38a4 4 0 0 1 4-4m-4 4a4 4 0 1 0 4-4m-4 4h-6m10-4v-6"/><path fill="#fff" d="M34 10a4 4 0 1 0 8 0a4 4 0 0 0-8 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 10a4 4 0 0 0 4 4m-4-4a4 4 0 1 1 4 4m-4-4h-6m10 4v6m-21 4h14m-7 7V17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLensAlignment0)"/>`),
		g.Group(children),
	)
}