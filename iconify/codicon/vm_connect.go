package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2h13l.5.5v5.503a5.006 5.006 0 0 0-1-.583V3H2v9h5a5 5 0 0 0 1 3H4v-1h3v-1H1.5l-.5-.5v-10l.5-.5z" clip-rule="evenodd"/><path d="M12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8zm0 7a3 3 0 1 1 0-6.001A3 3 0 0 1 12 15z"/><path fill-rule="evenodd" d="m12.133 11.435l1.436 1.436l.431-.431l-1.004-1.005L14 10.431l-.431-.43l-1.436 1.434zm-1.129 1.067L10 11.498l.431-.431l1.436 1.435l-1.436 1.436l-.431-.431l1.004-1.005z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}