package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highlighter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M398.789 22.31a31.762 31.762 0 0 0-22.771-9.52H376a31.769 31.769 0 0 0-22.783 9.552L87.534 292.234a32.086 32.086 0 0 0 .177 45.076l14.7 14.7L16 439.427V478h106.8l52.8-52.8l12.479 12.48a32 32 0 0 0 46-.77l258.234-276.14a31.913 31.913 0 0 0-.6-44.339ZM109.548 446H54.5l46.552-47.1l27.8 27.8Zm101.16-30.946L175.6 379.947l-24.127 24.127l-27.932-27.932l23.986-24.269l-37.191-37.189l48.338-49.105L257.8 364.7Zm68.958-73.74l-98.541-98.54L376.017 44.791l92.923 94.121Z"/>`),
		g.Group(children),
	)
}