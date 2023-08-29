package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoctailCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path d="M9 11a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1H9Z"/><path fill-rule="evenodd" d="m6.134 8.34l6.5 7a.5.5 0 0 0 .732 0l6.5-7a.5.5 0 0 0-.366-.84h-13a.5.5 0 0 0-.366.84Zm1.513.16h10.706L13 14.265L7.647 8.5Z" clip-rule="evenodd"/><path d="M12.5 14.875h1l.125.125v6l-.125.125h-1L12.375 21v-6l.125-.125Z"/><path d="M17.5 23h-9c0-1.475 2.05-2.5 4.5-2.5s4.5 1.025 4.5 2.5ZM15.879 4.567a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/><path d="M13.203 12.747a.5.5 0 0 1-.94-.343l3.093-8.493a.5.5 0 1 1 .94.342l-3.093 8.494Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}