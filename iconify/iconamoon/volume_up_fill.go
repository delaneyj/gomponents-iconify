package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.51 6.303c-.255-1.466-1.932-2.065-3.083-1.25L7.94 6.817A1 1 0 0 1 7.363 7H5a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h2.363a1 1 0 0 1 .578.184l2.486 1.762c1.15.816 2.828.217 3.083-1.25c.322-1.85.49-3.754.49-5.696c0-1.942-.168-3.845-.49-5.697Zm8.288 1.598a1 1 0 0 0-1.99.198a39.469 39.469 0 0 1 0 7.802a1 1 0 1 0 1.99.198a41.471 41.471 0 0 0 0-8.198Zm-3.925 1.017a1 1 0 1 0-1.993.164a35.506 35.506 0 0 1 0 5.836a1 1 0 0 0 1.993.164a37.465 37.465 0 0 0 0-6.164Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}