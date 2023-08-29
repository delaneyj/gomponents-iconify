package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m38.792 4l-3.614 36.439h-6.356L25.208 4h13.584M41 2H23l4.01 40.439h9.979L41 2zm-9 46.707c3.122 0 5.662 2.533 5.662 5.646S35.122 60 32 60c-3.123 0-5.664-2.533-5.664-5.646s2.541-5.647 5.664-5.647m0-2c-4.233 0-7.664 3.424-7.664 7.646S27.767 62 32 62c4.232 0 7.662-3.424 7.662-7.646s-3.43-7.647-7.662-7.647z"/>`),
		g.Group(children),
	)
}