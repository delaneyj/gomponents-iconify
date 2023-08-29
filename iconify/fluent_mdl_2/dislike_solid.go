package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 256q178 0 345 69q72 29 144 44t151 15h448v896h-417q-65 0-122 24t-104 70l-622 621q-25 25-50 39t-61 14q-33 0-62-12t-51-35t-34-51t-13-62q0-81 18-154t53-146q20-43 34-87t19-93H192q-39 0-74-15t-61-41t-42-61t-15-75q0-32 10-61l256-768q20-59 70-95t112-36h512z"/>`),
		g.Group(children),
	)
}