package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 4.48h-.71L2 3.43l.71-.7l.69.68L4.81 2l.71.71l-1.77 1.77zM6.99 3h8v1h-8V3zm0 3h8v1h-8V6zm8 3h-8v1h8V9zm-8 3h8v1h-8v-1zM3.04 7.48h.71l1.77-1.77l-.71-.7L3.4 6.42l-.69-.69l-.71.71l1.04 1.04zm.71 3.01h-.71L2 9.45l.71-.71l.69.69l1.41-1.42l.71.71l-1.77 1.77zm-.71 3.01h.71l1.77-1.77l-.71-.71l-1.41 1.42l-.69-.69l-.71.7l1.04 1.05z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}