package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.017.058v15.896h11.926V5.032H7.952V.058H2.017zm9.025 12.524c0 .367-.096.554-.288.554c-.205 0-.349-.255-.377-.306c-.519-.963-2.046-1.779-3.361-1.805v.905l-.039.038a.595.595 0 0 1-.429.155a.642.642 0 0 1-.393-.118l-2.108-1.603a.55.55 0 0 1 0-.756l2.097-1.613a.633.633 0 0 1 .447-.172c.158 0 .293.051.386.146l.039.039V9c1.873.026 4.026 1.582 4.026 3.582z"/><path d="M9.004.016v3.978h4.946L9.004.016z"/></g>`),
		g.Group(children),
	)
}