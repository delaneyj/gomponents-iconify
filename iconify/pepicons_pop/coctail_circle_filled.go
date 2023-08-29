package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoctailCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCoctailCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M9 11a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1H9Z"/><path fill-rule="evenodd" d="m5.767 8.68l6.5 7a1 1 0 0 0 1.466 0l6.5-7c.594-.64.14-1.68-.733-1.68h-13c-.873 0-1.327 1.04-.733 1.68Zm11.44.32L13 13.53L8.793 9h8.414Z" clip-rule="evenodd"/><path d="M14 14.5v7h-2v-7h2Z"/><path d="M17.5 23h-9c0-1.475 2.05-2.5 4.5-2.5s4.5 1.025 4.5 2.5ZM15.818 4.728a.75.75 0 0 1 .364-1.456l4 1a.75.75 0 0 1-.364 1.456l-4-1Z"/><path d="M14.211 11.737a.75.75 0 1 1-1.423-.474l2.5-7.5a.75.75 0 1 1 1.424.474l-2.5 7.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCoctailCircleFilled0)"/></g>`),
		g.Group(children),
	)
}