package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnDeviceTrainingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17v-1h2v1h-2Zm0-1.5v-.775q-.475-.275-.738-.738T10 13q0-.825.588-1.413T12 11q.825 0 1.413.588T14 13q0 .525-.263.988t-.737.737v.775h-2Zm5 .45l-1.05-1.075q.275-.425.413-.9T15.5 13q0-.5-.138-.975t-.412-.9L16 10.05q.5.65.75 1.4T17 13q0 .8-.25 1.55t-.75 1.4Zm-8 0q-.5-.65-.75-1.4T7 13q0-2.075 1.463-3.538T12 8V6.75l2.25 2l-2.25 2V9.5q-1.45 0-2.475 1.025T8.5 13q0 .5.138.975t.412.9L8 15.95ZM4 23V1h16v22H4Zm2-5h12V6H6v12Z"/>`),
		g.Group(children),
	)
}