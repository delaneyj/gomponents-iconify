package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h14v-1h7V7h-7V6H3Zm14 3v5h5V9h-5Zm-2 6V8H3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h12Zm4-3.992h1.01v-1.01H19v1.01Zm.51 2h-.5v-1.01h1.01v1.01h-.51Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}