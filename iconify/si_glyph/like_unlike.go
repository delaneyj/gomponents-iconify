package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LikeUnlike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.811 2.035H2.194c-.632 0-1.146.502-1.146 1.123v5.707c0 .621.515 1.123 1.146 1.123h2.617c.634 0 1.148-.502 1.148-1.123V3.158c0-.621-.515-1.123-1.148-1.123zm10.127 3.01V3.961h1.514c-.076-1.146-.658-1.898-1.74-1.898h-4.426c-.688 0-1.029 1.312-2.699 1.312l-.558-.018v5.955s1.063.166 1.419 1.291c0 0 1.451 3.961 2.57 4.357c.658 0 1.191.047 1.191-1.125l-.595-1.814s-.353-2.033 2.14-2.033h2.145c.688 0 1.06-.424 1.06-1.049c0 0 .014-.357.007-.895h-2.027V6.96h1.99a16.57 16.57 0 0 0-.218-1.916h-1.773v.001z"/>`),
		g.Group(children),
	)
}