package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeGrain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g stroke-width=".406"><path fill="#faa520" d="m16.056 2.024l-4.726 9.765l4.799 4.182l4.657-4.182z"/><path fill="#f78f28" d="m3.213 6.863l-.367 11.773l13.28 11.34l6.664-5.834z"/><path fill="#e18026" d="M29.153 18.64V6.726l-11.9 10.507l6.843 5.798z"/></g>`),
		g.Group(children),
	)
}