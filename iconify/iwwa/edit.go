package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M6.079 35.481a.48.48 0 0 1-.478-.492l.338-11.84a.477.477 0 0 1 .141-.324L26.152 2.75a.49.49 0 0 1 .676 0l10.826 10.826a.477.477 0 0 1 0 .677L17.581 34.327a.48.48 0 0 1-.295.139L6.122 35.479l-.043.002zm.811-12.114l-.317 11.11l10.455-.949L36.64 13.915L26.49 3.766L6.89 23.367z"/><path fill="currentColor" d="M32.029 19.466a.462.462 0 0 1-.33-.137L21.152 8.781a.465.465 0 0 1 0-.66a.465.465 0 0 1 .66 0l10.547 10.548a.465.465 0 0 1-.33.797z"/>`),
		g.Group(children),
	)
}