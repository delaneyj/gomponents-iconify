package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 4h-2C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.172 6.828c.943.944 2.348 1.127 4.828 1.163M16 4.01c2.48.036 3.885.22 4.828 1.163C22 6.343 22 8.229 22 12s0 5.657-1.172 6.828C19.657 20 17.771 20 14 20h-2m3-3H9"/>`),
		g.Group(children),
	)
}