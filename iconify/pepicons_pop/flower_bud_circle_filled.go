package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerBudCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFlowerBudCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M19.49 16.713a4.003 4.003 0 0 0 0-7.426A4 4 0 0 0 13 5.877a4 4 0 0 0-6.49 3.41a4.003 4.003 0 0 0 0 7.427a4 4 0 0 0 6.49 3.41a4 4 0 0 0 6.49-3.41Zm-2.108-7.035a1 1 0 0 0 .835 1.334a2 2 0 0 1 0 3.976a1 1 0 0 0-.835 1.334a2 2 0 0 1-3.55 1.783a1 1 0 0 0-1.665 0a2 2 0 0 1-3.55-1.783a1 1 0 0 0-.834-1.333a2 2 0 0 1 0-3.977a1 1 0 0 0 .835-1.334a2 2 0 0 1 3.55-1.783a1 1 0 0 0 1.665 0a2 2 0 0 1 3.55 1.783Z" clip-rule="evenodd"/><path d="M16.5 13a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFlowerBudCircleFilled0)"/></g>`),
		g.Group(children),
	)
}