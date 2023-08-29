package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.9 19.142l-.708-.707c-2.12-2.121-2.12-4.95 0-7.071l2.829-2.828c2.121-2.122 4.95-2.122 7.07 0c2.122 2.12 2.122 4.95 0 7.07m-6.363-2.12l.707.706c2.121 2.122 2.121 4.95 0 7.072l-2.828 2.828c-2.122 2.121-4.95 2.121-7.071 0c-2.122-2.121-2.122-4.95 0-7.071"/>`),
		g.Group(children),
	)
}