package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleOLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1152 672v192q0 13-9.5 22.5T1120 896H768v192q0 14-9 23t-23 9q-12 0-24-10L393 791q-9-9-9-23t9-23l320-320q9-9 23-9q13 0 22.5 9.5T768 448v192h352q13 0 22.5 9.5t9.5 22.5zm160 96q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273zm224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}