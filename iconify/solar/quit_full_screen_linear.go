package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuitFullScreenLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M14 22c0-3.771 0-5.657 1.172-6.828C16.343 14 18.229 14 22 14M2 14c3.771 0 5.657 0 6.828 1.172C10 16.343 10 18.229 10 22M2 10c3.771 0 5.657 0 6.828-1.172C10 7.657 10 5.771 10 2m12 8c-3.771 0-5.657 0-6.828-1.172C14 7.657 14 5.771 14 2"/>`),
		g.Group(children),
	)
}