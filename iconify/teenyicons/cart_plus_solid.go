package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartPlusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.021.644L.979.356L1.472 2H12.5A2.5 2.5 0 0 1 15 4.5V11H3.128L.02.644ZM8 9V7H6V6h2V4h1v2h2v1H9v2H8Zm-4 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm5.5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}