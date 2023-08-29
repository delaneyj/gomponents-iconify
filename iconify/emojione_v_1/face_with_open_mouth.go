package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fbbf67" d="M63.957 32.021C63.957 49.683 49.643 64 31.978 64S0 49.684 0 32.021C0 14.358 14.314.041 31.979.041s31.978 14.318 31.978 31.98"/><path fill="#fff" d="M27.552 21.234c0 5.38-2.907 9.732-6.498 9.732c-3.59 0-6.501-4.353-6.501-9.732c0-5.376 2.911-9.733 6.501-9.733c3.591 0 6.498 4.357 6.498 9.733m21.993 0c0 5.38-2.907 9.732-6.498 9.732c-3.59 0-6.501-4.353-6.501-9.732c0-5.376 2.911-9.733 6.501-9.733c3.591 0 6.498 4.357 6.498 9.733"/><g fill="#25333a"><ellipse cx="21.05" cy="21.234" rx="4.206" ry="5.099"/><ellipse cx="43.05" cy="21.234" rx="4.206" ry="5.099"/></g><ellipse cx="32.733" cy="46.908" fill="#633d19" rx="10.172" ry="10.09"/>`),
		g.Group(children),
	)
}