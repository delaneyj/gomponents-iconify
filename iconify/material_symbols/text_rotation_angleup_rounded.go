package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationAngleupRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.775 21.025q-.3-.3-.3-.712t.3-.713l8.625-8.625h-.6q-.425 0-.713-.275T15.8 10q0-.425.288-.713T16.8 9h3q.425 0 .725.3t.3.725V13q0 .425-.287.713t-.713.287q-.425 0-.725-.288T18.8 13v-.6l-8.625 8.625q-.3.3-.7.3t-.7-.3Zm.175-5.85q-.375.375-.8.3T7.5 14.9L3.575 6.3q-.175-.375-.088-.837t.388-.763q.275-.275.75-.375t.85.075l8.6 3.925q.5.225.575.65t-.325.825q-.175.175-.45.225t-.5-.075L11.25 8.9L8.1 12.05l1.025 2.15q.125.225.075.5t-.25.475Zm-1.575-4.65L9.7 8.2L5.35 6.125l-.05.05l2.075 4.35Z"/>`),
		g.Group(children),
	)
}