package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NormalWeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 1313q0 115-44 204t-118 149t-173 92t-205 32H512V128h473q88 0 172 22t150 68t106 118t40 172q0 69-19 132t-57 114t-91 90t-121 58v5q84 9 152 41t117 85t75 124t27 156zM707 304v536h198q72 0 134-16t109-52t73-91t27-134q0-72-26-119t-70-74t-101-39t-120-11H707zm625 1006q0-90-34-147t-91-89t-129-45t-149-12H707v597h263q73 0 138-17t115-54t79-94t30-139z"/>`),
		g.Group(children),
	)
}