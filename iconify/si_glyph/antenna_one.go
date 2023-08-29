package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.987 0H3.112c-2.127 0-.266 1.982-.266 1.982l5.234 6.08v7.854h1.875V8.077l5.232-6.064c.001 0 1.955-2.013-.2-2.013zM7.64 5.582L4.204 1.52S3.526.97 4.35.97c.825 0 3.668.014 3.668.014v4.531c0 .673-.378.067-.378.067zm2.302-.066V.985s2.909-.014 3.752-.014c.844 0 .15.55.15.55L10.33 5.583c.001-.001-.388.605-.388-.067z"/>`),
		g.Group(children),
	)
}