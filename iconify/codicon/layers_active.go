package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.185 1.083l-.558.004l-5.909 4.037l.004.828L7.63 9.915l.55.004l6.092-3.963l.003-.836l-6.09-4.037Zm-5.293 4.45l5.021-3.43l5.176 3.43l-5.176 3.368l-5.021-3.368Zm4.739 6.882L1.793 8.5h1.795l4.325 2.9l4.459-2.9h1.833l-.8.52a4.002 4.002 0 0 0-4.21 2.739l-1.013.66l-.551-.004Zm1.373.776l-1.09.71L3.587 11H1.793l5.838 3.915l.55.004l1.02-.663a3.988 3.988 0 0 1-.197-1.065Zm2.329-2.685a3 3 0 1 1 3.333 4.987a3 3 0 0 1-3.333-4.987Zm1.698 3.817l1.79-2.387l-.8-.6l-1.48 1.974l-.876-.7l-.624.78l1.278 1.023l.712-.09Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}