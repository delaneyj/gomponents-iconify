package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightBulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 3.991C13 1.422 10.729.015 8 .015S3 1.421 3 3.991c0 3.299 2.087 4.197 3.312 6.278c.264.449-.49.227-.235.683h3.818c.252-.452-.595-.229-.333-.676C10.782 8.192 13 7.375 13 3.991zM6 12h3.953v.922H6zm3.969 2.968c0 .561-.434 1.017-.971 1.017H7.014c-.538 0-.972-.456-.972-1.017v-.947h3.927v.947z"/>`),
		g.Group(children),
	)
}