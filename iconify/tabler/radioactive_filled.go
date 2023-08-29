package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioactiveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M21 11a1 1 0 0 1 1 1a10 10 0 0 1-5 8.656a1 1 0 0 1-1.302-.268l-.064-.098l-3-5.19a.995.995 0 0 1-.133-.542l.01-.11l.023-.106l.034-.106l.046-.1l.056-.094l.067-.089a.994.994 0 0 1 .165-.155l.098-.064a2 2 0 0 0 .993-1.57l.007-.163a1 1 0 0 1 .883-.994L15 11h6zM7 3.344a10 10 0 0 1 10 0a1 1 0 0 1 .418 1.262l-.052.104l-3 5.19l-.064.098a.994.994 0 0 1-.155.165l-.089.067a1 1 0 0 1-.195.102l-.105.034l-.107.022a1.003 1.003 0 0 1-.547-.07L13 10.266a2 2 0 0 0-1.842-.082l-.158.082a1 1 0 0 1-1.302-.268L9.634 9.9l-3-5.19A1 1 0 0 1 7 3.344zM9 11a1 1 0 0 1 .993.884l.007.117a2 2 0 0 0 .861 1.645l.237.152a.994.994 0 0 1 .165.155l.067.089l.056.095l.045.099c.014.036.026.07.035.106l.022.107l.011.11a.994.994 0 0 1-.08.437l-.053.104l-3 5.19A1 1 0 0 1 7 20.656A10 10 0 0 1 2 12a1 1 0 0 1 .883-.993L3 11h6z"/></g>`),
		g.Group(children),
	)
}