package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Solidity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M82.26 292.547L0 146.302L82.26 0h164.574L330 146.302H164.518L82.26 292.547zM330 365.697l-82.369-146.244l-82.26 146.245H0L83.056 512H247.63L330 365.697z"/>`),
		g.Group(children),
	)
}