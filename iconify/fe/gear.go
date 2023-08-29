package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feGear0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGear1" fill="currentColor" fill-rule="nonzero"><path id="feGear2" d="M11 19.938a7.96 7.96 0 0 1-3.906-1.618l-1.458 1.458l-1.414-1.414l1.458-1.458A7.96 7.96 0 0 1 4.062 13H2v-2h2.062A7.96 7.96 0 0 1 5.68 7.094L4.222 5.636l1.414-1.414L7.094 5.68A7.96 7.96 0 0 1 11 4.062V2h2v2.062a7.96 7.96 0 0 1 3.906 1.618l1.458-1.458l1.414 1.414l-1.458 1.458A7.96 7.96 0 0 1 19.938 11H22v2h-2.062a7.96 7.96 0 0 1-1.618 3.906l1.458 1.458l-1.414 1.414l-1.458-1.458A7.96 7.96 0 0 1 13 19.938V22h-2v-2.062ZM12 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></g>`),
		g.Group(children),
	)
}