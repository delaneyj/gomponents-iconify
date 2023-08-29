package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Confluence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.864 29.145L4.64 34.728c-.288.5-.117 1.138.382 1.426l6.859 3.96c.5.288 1.137.117 1.426-.382l2.394-4.147a4.175 4.175 0 0 1 5.703-1.529l10.178 5.877c.5.288 1.138.117 1.426-.382l3.96-6.859a1.045 1.045 0 0 0-.382-1.426l-11.614-6.705c-5.99-3.458-13.65-1.406-17.108 4.584Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.136 18.855l3.224-5.583c.288-.5.117-1.138-.382-1.426l-6.859-3.96a1.044 1.044 0 0 0-1.426.382L32.3 12.415a4.175 4.175 0 0 1-5.703 1.529L16.418 8.067a1.044 1.044 0 0 0-1.426.382l-3.96 6.859a1.045 1.045 0 0 0 .382 1.426l11.614 6.705c5.99 3.459 13.65 1.406 17.108-4.584Z"/>`),
		g.Group(children),
	)
}