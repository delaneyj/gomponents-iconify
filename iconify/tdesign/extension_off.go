package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3.99A3 3 0 0 1 14 4h6v6a3 3 0 0 1 .01 6l.05 2.516L18 16.408V14h2a1 1 0 1 0 0-2h-2V6h-6V4a1 1 0 1 0-2 0v2H7.586L5.528 3.942L8 3.99ZM1 1.586L22.414 23L21 24.414L18.586 22h-4.988l-.123-.858a2.5 2.5 0 0 0-4.95 0L8.404 22H2v-6.403l.859-.122a2.5 2.5 0 0 0 0-4.95L2 10.403V5.414L-.414 3L1 1.586Zm3 5.828v1.342a4.501 4.501 0 0 1 0 8.488V20h2.756a4.501 4.501 0 0 1 8.488 0h1.342L4 7.414Z"/>`),
		g.Group(children),
	)
}