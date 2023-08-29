package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 15v-1c2 0 2-.6 2-1H1.5l-.5-.5v-10l.5-.5h13l.5.5v9.24l-1-1V3H2v9h5.73l-.5.5l2.5 2.5H4zm7.86 0l2.5-2.5l-.71-.7L12 13.45V7h-1v6.44l-1.64-1.65l-.71.71l2.5 2.5h.71z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}