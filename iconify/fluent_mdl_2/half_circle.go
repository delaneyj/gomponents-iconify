package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalfCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 0v1920q-59 0-114-3t-109-11t-107-23t-108-38q-134-57-242-147t-184-206t-118-251t-42-281q0-133 34-255t96-230t150-194t195-150t229-97t256-34h64zm-128 130q-108 8-207 42t-184 91t-155 131t-119 165t-76 191t-27 210q0 108 27 209t76 191t119 165t155 132t184 90t207 43V130z"/>`),
		g.Group(children),
	)
}