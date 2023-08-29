package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMaldives(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M64 43c0 6.075-3.373 11-10 11H10C3.373 54 0 49.075 0 43V21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v22"/><path fill="#137a08" d="M9 18h47v26H9z"/><path fill="#e6e7e8" fill-rule="evenodd" d="M37.654 40.738a10.025 10.025 0 0 1-10.412-3.753c-2.632-3.494-2.666-8.406-.111-11.907a10.059 10.059 0 0 1 10.535-3.905c-2.711.138-4.996 1.12-6.832 3.059c-1.836 1.939-2.73 4.26-2.689 6.923c.076 5.078 4.156 9.39 9.509 9.583"/>`),
		g.Group(children),
	)
}