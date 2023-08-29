package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TzFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagTz4x30"><path fill-opacity=".7" d="M10 0h160v120H10z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagTz4x30)" transform="matrix(4 0 0 4 -40 0)"><path fill="#09f" d="M0 0h180v120H0z"/><path fill="#090" d="M0 0h180L0 120V0z"/><path d="M0 120h40l140-95V0h-40L0 95v25z"/><path fill="#ff0" d="M0 91.5L137.2 0h13.5L0 100.5v-9zM29.3 120L180 19.5v9L42.8 120H29.3z"/></g>`),
		g.Group(children),
	)
}