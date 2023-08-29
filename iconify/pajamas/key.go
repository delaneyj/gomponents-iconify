package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.358 8.763l.675-.674l-.325-.897A3.5 3.5 0 1 1 10 9.5H7.5v1H6.379l-.44.44l-1 1l-.439.439V13.5h-2v-.879l3.858-3.858ZM6 15v-2l1-1h2v-1h1a5 5 0 1 0-4.703-3.297L1 12v3h5Zm5-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}