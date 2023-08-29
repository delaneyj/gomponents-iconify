package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Google(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 7v2.4h3.97c-.16 1.03-1.2 3.02-3.97 3.02c-2.39 0-4.34-1.98-4.34-4.42S5.61 3.58 8 3.58c1.36 0 2.27.58 2.79 1.08l1.9-1.83C11.47 1.69 9.89 1 8 1C4.13 1 1 4.13 1 8s3.13 7 7 7c4.04 0 6.72-2.84 6.72-6.84c0-.46-.05-.81-.11-1.16H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}