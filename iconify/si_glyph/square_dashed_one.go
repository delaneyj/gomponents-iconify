package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareDashedOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.048 1.976h1.993v1.988h.907v-2.92h-2.9v.932zM.923 3.995V1.976h2.028v-.99H.027v3.009h.896zm2.028 12.022H.964v-1.95H.027v2.928h2.924v-.978zM15 6v2.01h.989V6H15zm0 4v1.96h.982V10H15zM9 1v.965h1.984V1H9zM5 1v.973h1.995V1H5zm4 15v.973h1.958V16H9zm-4 0v.993h1.987V16H5zM0 6v1.99h.98V6H0zm0 4v1.961h.954V10H0zm15.006 4v2.073H13v.945h3.003V14h-.997z"/>`),
		g.Group(children),
	)
}