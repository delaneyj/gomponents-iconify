package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.09 21.5a.67.67 0 0 1-.24 0a.83.83 0 0 1-.59-.81v-.11l.9-6.35H6.82a.8.8 0 0 1-.71-.43a.85.85 0 0 1 0-.86l2-3.49l4.1-6.52a.79.79 0 0 1 .92-.35a.83.83 0 0 1 .59.81v.11l-.9 6.35h4.35a.8.8 0 0 1 .71.43a.85.85 0 0 1 0 .86l-2 3.49l-4.1 6.52a.79.79 0 0 1-.69.35Zm-.53-1.21Zm-2.63-7.6h4a.84.84 0 0 1 .83.85v.11l-.59 4.14l2.5-4l1.44-2.48h-4a.84.84 0 0 1-.83-.85v-.11l.59-4.14l-2.5 4Zm-.57 1Zm9.28-3.34Zm-3.2-6.65Z"/>`),
		g.Group(children),
	)
}