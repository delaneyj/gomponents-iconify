package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSimpleSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.94 210.62a8 8 0 1 1-11.84 10.76L182.66 200H48a16 16 0 0 1-13.8-24.06C39.75 166.38 48 139.34 48 104a79.49 79.49 0 0 1 10.84-40.2L42.1 45.38a8 8 0 1 1 11.84-10.76ZM160 216H96a8 8 0 1 0 0 16h64a8 8 0 0 0 0-16Zm58.51-47.12C216.07 162.59 208 139.35 208 104A80 80 0 0 0 88.54 34.4a8 8 0 0 0-2 12.34l118.63 130.44a8 8 0 0 0 13.36-8.3Z"/>`),
		g.Group(children),
	)
}