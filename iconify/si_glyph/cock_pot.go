package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CockPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.127 11.941H1.883C.529 11.941.016 7.004.016 7.004h9.979c0 .001-.268 4.937-1.868 4.937z"/><path d="M15.984 7.486c0 .252-.276.456-.615.456H8.807c-.34 0-.614-.204-.614-.456s.274-.457.614-.457h6.562c.339 0 .615.205.615.457zM4 4h1.922v1.297H4z"/><path d="M0 5h9.906v.969H0z"/></g>`),
		g.Group(children),
	)
}