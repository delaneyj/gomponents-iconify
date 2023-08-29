package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerClawCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopHammerClawCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.632 4.703a2 2 0 0 1 1.088.179l3.808 1.802a2 2 0 0 1 .952 2.664l-1.07 2.26a2 2 0 0 1-2.663.952l-3.342-1.582a2 2 0 0 1-.859-.778l-.35-.583l-.95.339c-1.677.6-3.243-1.128-2.48-2.739l.44-.932a2 2 0 0 1 1.576-1.131l3.85-.45Zm4.04 3.79L11.864 6.69l-3.85.45l-.44.933l.949-.34a2 2 0 0 1 2.388.854l.35.583l3.342 1.582l1.07-2.26Z"/><path d="M16.79 8.468a2 2 0 0 1 2.663-.952l.904.428a2 2 0 0 1 .952 2.663l-1.71 3.616a2 2 0 0 1-2.664.952l-.904-.428a2 2 0 0 1-.952-2.663l1.711-3.616Zm2.712 1.284l-.904-.428l-1.712 3.615l.904.428l1.712-3.615Zm-7.349 1.5l-3.85 8.135l.904.428l3.85-8.135l1.808.856l-3.85 8.134a2 2 0 0 1-2.664.952l-.904-.427a2 2 0 0 1-.952-2.664l3.85-8.135l1.808.856Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopHammerClawCircleFilled0)"/></g>`),
		g.Group(children),
	)
}