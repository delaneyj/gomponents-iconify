package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.017.058v15.896h12.006V5.032H7.952V.058H2.017zM8.922 9v-.953l.039-.039a.527.527 0 0 1 .386-.146c.174 0 .341.063.447.172l2.097 1.613a.55.55 0 0 1 0 .756l-2.108 1.603a.64.64 0 0 1-.393.118a.6.6 0 0 1-.429-.155l-.039-.038v-.905c-1.315.025-2.843.842-3.361 1.805c-.029.051-.173.306-.377.306c-.192 0-.288-.187-.288-.554c0-2.001 2.153-3.557 4.026-3.583z"/><path d="M9.004.016v3.978h4.946L9.004.016z"/></g>`),
		g.Group(children),
	)
}