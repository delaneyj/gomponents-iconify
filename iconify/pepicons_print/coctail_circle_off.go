package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoctailCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M15.75 21.024v-5.097l6.116-6.587a.5.5 0 0 0-.366-.84h-13a.5.5 0 0 0-.366.84l6.116 6.587v5.097c-2.117.139-3.75.866-3.75 1.976h9c0-1.11-1.633-1.837-3.75-1.976Z" clip-rule="evenodd" opacity=".2"/><path d="M9 11a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1H9Z"/><path fill-rule="evenodd" d="m6.134 8.34l6.5 7a.5.5 0 0 0 .732 0l6.5-7a.5.5 0 0 0-.366-.84h-13a.5.5 0 0 0-.366.84Zm1.513.16h10.706L13 14.265L7.647 8.5Z" clip-rule="evenodd"/><path d="M12.5 14.875h1l.125.125v6l-.125.125h-1L12.375 21v-6l.125-.125Z"/><path d="M17.5 23h-9c0-1.475 2.05-2.5 4.5-2.5s4.5 1.025 4.5 2.5ZM15.879 4.567a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/><path d="M13.203 12.747a.5.5 0 0 1-.94-.343l3.093-8.493a.5.5 0 1 1 .94.342l-3.093 8.494Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}