package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmRunning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2h13l.5.5v5.503a5.006 5.006 0 0 0-1-.583V3H2v9h5a5 5 0 0 0 1 3H4v-1h3v-1H1.5l-.5-.5v-10l.5-.5z" clip-rule="evenodd"/><path d="M12 8c.367 0 .721.047 1.063.14c.34.094.658.23.953.407c.294.177.563.385.808.625c.245.24.455.509.63.808a4.03 4.03 0 0 1 .405 3.082c-.093.342-.229.66-.406.954a4.382 4.382 0 0 1-.625.808c-.24.245-.509.455-.808.63a4.029 4.029 0 0 1-3.082.405a3.784 3.784 0 0 1-.954-.406a4.382 4.382 0 0 1-.808-.625a3.808 3.808 0 0 1-.63-.808a4.027 4.027 0 0 1-.405-3.082c.093-.342.229-.66.406-.954c.177-.294.385-.563.625-.808c.24-.245.509-.455.808-.63A4.028 4.028 0 0 1 12 8zm2 3.988L11 10v4l3-2.012z"/></g>`),
		g.Group(children),
	)
}