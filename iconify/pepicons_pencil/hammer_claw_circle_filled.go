package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerClawCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilHammerClawCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.69 5.2a1.5 1.5 0 0 1 .816.134l3.808 1.802a1.5 1.5 0 0 1 .714 1.998l-1.07 2.26a1.5 1.5 0 0 1-1.997.714l-3.342-1.582a1.5 1.5 0 0 1-.644-.584l-.35-.583a.5.5 0 0 0-.597-.213l-.95.34c-1.258.449-2.432-.847-1.86-2.055l.44-.932A1.5 1.5 0 0 1 7.84 5.65l3.85-.45Zm.388 1.038a.5.5 0 0 0-.272-.045l-3.85.45a.5.5 0 0 0-.393.283l-.442.933a.5.5 0 0 0 .62.684l.95-.339a1.5 1.5 0 0 1 1.791.64l.35.583a.5.5 0 0 0 .215.195l3.342 1.582a.5.5 0 0 0 .666-.238l1.07-2.26a.5.5 0 0 0-.239-.666l-3.808-1.802Z"/><path d="M17.242 8.682a1.5 1.5 0 0 1 1.997-.714l.904.428a1.5 1.5 0 0 1 .715 1.997l-1.642 3.467a1.5 1.5 0 0 1-1.997.714l-.904-.428a1.5 1.5 0 0 1-.714-1.997l1.641-3.467Zm1.57.19a.5.5 0 0 0-.666.238l-1.641 3.467a.5.5 0 0 0 .238.665l.904.428a.5.5 0 0 0 .666-.238l1.64-3.467a.5.5 0 0 0-.237-.665l-.904-.428ZM9.658 20.027l3.89-8.24l.904.426l-3.89 8.241a1.5 1.5 0 0 1-1.998.716l-.903-.427a1.5 1.5 0 0 1-.714-1.998l3.85-8.135l.904.428l-3.85 8.135a.5.5 0 0 0 .238.666l.903.427a.5.5 0 0 0 .666-.238Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilHammerClawCircleFilled0)"/></g>`),
		g.Group(children),
	)
}