package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.427 5.054c1.15-.816 2.828-.217 3.083 1.25c.322 1.85.49 3.754.49 5.696c0 1.942-.168 3.845-.49 5.697c-.255 1.466-1.932 2.065-3.083 1.25L7.94 17.183A1 1 0 0 0 7.363 17H5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3h2.363a1 1 0 0 0 .578-.184l2.486-1.762Zm7.28 4.239a1 1 0 0 0-1.414 1.414L17.586 12l-1.293 1.293a1 1 0 0 0 1.414 1.414L19 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L20.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L19 10.586l-1.293-1.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}