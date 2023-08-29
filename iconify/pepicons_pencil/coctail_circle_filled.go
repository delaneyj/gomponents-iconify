package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoctailCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCoctailCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M9 11a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1H9Z"/><path fill-rule="evenodd" d="m6.134 8.34l6.5 7a.5.5 0 0 0 .732 0l6.5-7a.5.5 0 0 0-.366-.84h-13a.5.5 0 0 0-.366.84Zm1.513.16h10.706L13 14.265L7.647 8.5Z" clip-rule="evenodd"/><path d="M12.5 14.875h1l.125.125v6l-.125.125h-1L12.375 21v-6l.125-.125Z"/><path d="M17.5 23h-9c0-1.475 2.05-2.5 4.5-2.5s4.5 1.025 4.5 2.5ZM15.879 4.567a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/><path d="M13.203 12.747a.5.5 0 0 1-.94-.343l3.093-8.493a.5.5 0 1 1 .94.342l-3.093 8.494Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCoctailCircleFilled0)"/></g>`),
		g.Group(children),
	)
}