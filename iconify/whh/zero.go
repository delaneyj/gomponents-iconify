package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024H256q-106 0-181-75T0 768V256Q0 150 75 75T256 0h256q106 0 181 75t75 181v512q0 106-75 181t-181 75zM256 896h256q31 0 61-16L128 267v501q0 53 37.5 90.5T256 896zm256-768H256q-32 0-61 16l445 613V256q0-53-37.5-90.5T512 128z"/>`),
		g.Group(children),
	)
}