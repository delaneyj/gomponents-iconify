package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<defs><path id="pepiconsClipboard0" fill="currentColor" d="M16.086 18H3.907A1.107 1.107 0 0 1 2.8 16.893V3.607c0-.611.496-1.107 1.107-1.107H7.23v2.214H5.014v11.072h9.965V4.714h-2.215V2.5h3.322c.611 0 1.107.496 1.107 1.107v13.286c0 .611-.496 1.107-1.107 1.107Z"/></defs><g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6.675 2.5h6.643v3.875H6.675z"/><use href="#pepiconsClipboard0"/><path fill="currentColor" fill-rule="evenodd" d="M5.675 2.5a1 1 0 0 1 1-1h6.643a1 1 0 0 1 1 1v3.875a1 1 0 0 1-1 1H6.675a1 1 0 0 1-1-1V2.5Zm2 1v1.875h4.643V3.5H7.675Z" clip-rule="evenodd"/><use href="#pepiconsClipboard0"/></g>`),
		g.Group(children),
	)
}