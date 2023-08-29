package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightRuler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="m13.563 18.734l34.629 42.242l10.584-8.17l-34.723-42.613l-10.49 8.54z"/><path fill="#9B9B9A" d="m23.888 10.428l-2.601 2.03l34.561 42.679l2.42-2.547z"/><path fill="#d0cfce" d="m19.727 20.938l1.552-1.262"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m58.648 52.752l-10.642 8.605l-34.54-42.714l10.642-8.606zM47.49 55.073l1.551-1.262m-4.075-1.841l1.551-1.262m-4.075-1.841l1.551-1.262m-6.599-4.945l1.552-1.262m-4.076-1.841l1.552-1.262m-4.076-1.841l1.552-1.262m-4.075-1.841l1.551-1.262m-6.599-4.945l1.551-1.262m-4.075-1.841l1.551-1.262m-4.075-1.841l1.552-1.262M39.812 45.85l2.648-2.154M27.193 30.334l2.647-2.154"/>`),
		g.Group(children),
	)
}