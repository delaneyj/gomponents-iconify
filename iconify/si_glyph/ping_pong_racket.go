package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PingPongRacket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.297 10.802c-.451-.452-.447-.899-.372-1.144c.2-.67.474-1.415.632-2.234L8.53 12.451c.793-.162 1.517-.428 2.168-.621c.25-.074.723-.053 1.197.421c.473.475 2.906 3.642 2.906 3.642l2.164-2.163c0-.001-3.216-2.476-3.668-2.928zm.021-4.636c-.01-1.285-.493-2.733-2.094-4.336a5.938 5.938 0 0 0-8.396 8.396c1.484 1.483 2.832 1.993 4.047 2.062l6.443-6.122zm-6.847 2.9a1.537 1.537 0 0 1-1.526-1.549c0-.854.684-1.547 1.526-1.547c.844 0 1.528.693 1.528 1.547c0 .857-.685 1.549-1.528 1.549z"/>`),
		g.Group(children),
	)
}