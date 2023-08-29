package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerClaw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.632 1.703a2 2 0 0 1 1.088.179l3.808 1.802a2 2 0 0 1 .952 2.664l-1.07 2.26a2 2 0 0 1-2.663.952L7.405 7.978a2 2 0 0 1-.859-.778l-.35-.583l-.95.339c-1.677.6-3.243-1.128-2.48-2.739l.44-.932a2 2 0 0 1 1.576-1.131l3.85-.45Zm4.04 3.79L8.864 3.69l-3.85.45l-.44.933l.949-.34a2 2 0 0 1 2.388.854l.35.583l3.342 1.582l1.07-2.26Z"/><path d="M13.79 5.468a2 2 0 0 1 2.663-.952l.904.428a2 2 0 0 1 .952 2.663l-1.71 3.616a2 2 0 0 1-2.664.952l-.904-.428a2 2 0 0 1-.952-2.663l1.711-3.616Zm2.712 1.284l-.904-.428l-1.712 3.615l.904.428l1.712-3.615Zm-7.349 1.5l-3.85 8.135l.904.428l3.85-8.135l1.808.856l-3.85 8.134a2 2 0 0 1-2.664.952l-.904-.427a2 2 0 0 1-.952-2.664l3.85-8.135l1.808.856Z"/></g>`),
		g.Group(children),
	)
}