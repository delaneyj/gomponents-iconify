package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.969 27l-4.25 2C7.531 28.356 2 24.531 2 19.887c0-4.469 5.156-8.188 11.981-9.019v2.688c-4.469.781-7.75 3.313-7.75 6.331c0 3.188 3.656 5.831 8.481 6.438V5.075L18.962 3v24zM30 18.188l-8.206-1.781l2.3-1.294c-1.219-.719-2.719-1.25-4.375-1.55v-2.688c2.887.344 5.481 1.219 7.519 2.456l2.188-1.238z"/>`),
		g.Group(children),
	)
}