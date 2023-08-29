package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Construction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.34 16.06a1 1 0 0 0-1.108.3l-3.772 4.526l-5.436-.988l-3.602-8.952A3.014 3.014 0 0 0 12.614 9h-4.06a3.002 3.002 0 0 0-1.544.428L2 12.434v6.4l5 .91V30h2v-9.893l3.565.648L14 24.2V30h2v-6.2l-1.091-2.618l8.081 1.469l-4.758 5.709A1 1 0 0 0 19 30h10a1 1 0 0 0 1-1V17a1 1 0 0 0-.66-.94zM4 17.166v-3.6l3-1.8v5.945zm5 .909V11h3.614a1.014 1.014 0 0 1 .945.67l3.14 7.805zM28 28h-6.865L28 19.762zM12.5 8A3.5 3.5 0 1 1 16 4.5A3.504 3.504 0 0 1 12.5 8zm0-5A1.5 1.5 0 1 0 14 4.5A1.502 1.502 0 0 0 12.5 3z"/>`),
		g.Group(children),
	)
}