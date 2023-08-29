package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XxFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" fill-rule="evenodd" stroke="#adb5bd" stroke-width="1.1" d="M.5.5h638.9v478.9H.5z"/><path fill="none" stroke="#adb5bd" stroke-width="1.1" d="m.5.5l639 479m0-479l-639 479"/>`),
		g.Group(children),
	)
}