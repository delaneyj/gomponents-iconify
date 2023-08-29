package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.667.037H3.332A1.32 1.32 0 0 0 2.016 1.36l.01 14.075c0 .376.642.58.995.58c.354 0 .938-.204.938-.58v-1.407c.009 0 .017.003.017.003h10.01c.013 0 .023-.003.036-.004v1.408c0 .376.635.602.979.602c.347 0 .958-.226.958-.602l.027-14.075A1.322 1.322 0 0 0 14.667.037zm-1.688.921h1.041V3h-1.041V.958zm-2.021 0h1v2h-1v-2zm-1.985 0h1.048v2H8.973v-2zm6.048 9.063H2.98V8.98h12.041v1.041zM3.958 8V6h1v2h-1zm2.014-.007v-2h1v2h-1zM7.961 8V5.993h1V8h-1zm7.06-3H2.98V3.979h12.041V5z"/>`),
		g.Group(children),
	)
}