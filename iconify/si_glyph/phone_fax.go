package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneFax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.964 13.945c0 .565-.476 1.023-1.063 1.023h-.807c-.586 0-1.062-.458-1.062-1.023V1.054c0-.564.477-1.023 1.062-1.023h.807c.588 0 1.063.459 1.063 1.023v12.891zM15.881 0H6.034A1.03 1.03 0 0 0 5 1.023v12.891c0 .565.463 1.023 1.034 1.023h.924v-2.021H15v2.021h.881c.573 0 1.035-.458 1.035-1.023V1.023C16.916.459 16.454 0 15.881 0zM8 12.021H6.969v-1.053H8v1.053zm0-2H6.969V8.968H8v1.053zm0-2H6.969V6.984H8v1.037zm2 4H8.969v-1.053H10v1.053zm0-2H8.969V8.953H10v1.068zm0-2H8.969V6.968H10v1.053zm2 4h-1.031v-1.037H12v1.037zm0-2h-1.031V8.984H12v1.037zm0-2h-1.031V6.953H12v1.068zm3.016.954v1.047h-2.047V8.975h2.047zm-2.047-.954V6.959h2.047v1.062h-2.047zm2.047-3.99H6.969V1.969h8.047v2.062z"/><path d="m13.969 15.969l-.002-1.947H8.018l.002 1.947h5.949z"/></g>`),
		g.Group(children),
	)
}