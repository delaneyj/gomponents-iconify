package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GemStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M54.5 9.454L41 4H23L9.5 9.454L2 20.057L32 60l30-39.943l-7.5-10.603zM5.865 18.057l4.929-6.968L20.05 7.35l-2.82 10.708H5.865zM32 57L18.77 20.057h25.71L32 57zM43.756 7.271l9.451 3.818l4.928 6.968H46.094L43.756 7.271z"/>`),
		g.Group(children),
	)
}