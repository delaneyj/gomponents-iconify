package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoctailOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M6 8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1H6Z"/><path fill-rule="evenodd" d="m3.134 5.34l6.5 7a.5.5 0 0 0 .732 0l6.5-7a.5.5 0 0 0-.366-.84h-13a.5.5 0 0 0-.366.84Zm1.513.16h10.706L10 11.265L4.647 5.5Z" clip-rule="evenodd"/><path d="M9.5 11.875h1l.125.125v6l-.125.125h-1L9.375 18v-6l.125-.125Z"/><path d="M14.5 20h-9c0-1.475 2.05-2.5 4.5-2.5s4.5 1.025 4.5 2.5ZM12.879 1.567a.5.5 0 0 1 .242-.97l4 1a.5.5 0 0 1-.242.97l-4-1Z"/><path d="M10.203 9.747a.5.5 0 0 1-.94-.343L12.356.911a.5.5 0 1 1 .94.342l-3.093 8.494Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}