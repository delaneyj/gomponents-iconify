package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m12.115 2.612l.525-.812l.85.554l.484-.75L11.527.011l-.483.75l.849.553l-.512.788A7.394 7.394 0 0 0 7.502.999C3.387.999.041 4.352.041 8.475c0 4.12 3.346 7.474 7.461 7.474c4.113 0 7.461-3.354 7.461-7.474a7.463 7.463 0 0 0-2.848-5.863zM7.502 14.011c-3.047 0-5.527-2.488-5.527-5.544c0-3.058 2.48-5.545 5.527-5.545s5.527 2.487 5.527 5.545c0 3.055-2.48 5.544-5.527 5.544z"/><path d="M7 4h1v5H7z"/><path d="M7 8h3v1H7z"/></g>`),
		g.Group(children),
	)
}