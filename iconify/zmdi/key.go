package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M249 149h220v86h-42v85h-86v-85h-92q-14 37-47 61t-74 24q-53 0-90.5-37.5T0 192t37.5-90.5T128 64q41 0 74 24t47 61zm-121 86q18 0 30.5-12.5T171 192t-12.5-30.5T128 149t-30.5 12.5T85 192t12.5 30.5T128 235z"/>`),
		g.Group(children),
	)
}