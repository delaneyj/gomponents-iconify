package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768t-103 385.5t-279.5 279.5T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0zm0 1312q148 0 273-73t198-198t73-273t-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73zm96-224q-14 0-23-9t-9-23V480q0-14 9-23t23-9h192q14 0 23 9t9 23v576q0 14-9 23t-23 9H864zm-384 0q-14 0-23-9t-9-23V480q0-14 9-23t23-9h192q14 0 23 9t9 23v576q0 14-9 23t-23 9H480z"/>`),
		g.Group(children),
	)
}