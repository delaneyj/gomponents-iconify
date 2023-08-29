package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3h9v9H3V3ZM2 3a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3Zm8.35 2.511a.5.5 0 0 0-.825-.566L6.64 9.15L5.197 7.41a.5.5 0 0 0-.77.638l1.866 2.25a.5.5 0 0 0 .797-.037l3.26-4.749Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}