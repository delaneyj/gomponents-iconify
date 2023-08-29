package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.615.161a.837.837 0 0 1 .36.862l.001-.005l-3.426 20.56a.848.848 0 0 1-.424.6l-.004.002a.819.819 0 0 1-.406.107h-.01a.9.9 0 0 1-.326-.069l.006.002l-7.054-2.88l-3.989 4.377a.79.79 0 0 1-.604.281h-.026h.001h-.026a.76.76 0 0 1-.287-.056l.005.002a.808.808 0 0 1-.398-.311l-.002-.003a.856.856 0 0 1-.147-.482v-6.055L.539 14.51a.787.787 0 0 1-.535-.736a.78.78 0 0 1 .422-.788l.004-.002L22.705.134a.772.772 0 0 1 .912.027L23.615.16zm-4.578 20.065l2.96-17.709l-19.196 11.07l4.498 1.834L18.85 6.868l-6.4 10.668z"/>`),
		g.Group(children),
	)
}