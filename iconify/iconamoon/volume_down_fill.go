package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.427 5.054c1.15-.816 2.828-.217 3.083 1.25c.322 1.85.49 3.754.49 5.696c0 1.942-.168 3.845-.49 5.697c-.255 1.466-1.932 2.065-3.083 1.25L7.94 17.183A1 1 0 0 0 7.363 17H5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3h2.363a1 1 0 0 0 .578-.184l2.486-1.762Zm7.446 3.864a1 1 0 1 0-1.993.164a35.506 35.506 0 0 1 0 5.836a1 1 0 0 0 1.993.164a37.465 37.465 0 0 0 0-6.164Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}