package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kashflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.278 2.141l-.83 2.702C8.007.174 2.958 4.724 2.958 4.724C-1.638 8.564.49 14.678.495 14.678C1.252-.016 14.24 8.943 14.24 8.943c-.237 1.066-.996 2.63-.972 2.654l8.508-1.256zm7.228 7.181C22.747 24.016 9.76 15.057 9.76 15.057c.332-1.066 1.02-2.654 1.02-2.607l-8.51 1.21l5.451 8.2l.83-2.702c7.441 4.669 12.49.119 12.49.119c4.597-3.84 2.464-9.954 2.464-9.954z"/>`),
		g.Group(children),
	)
}