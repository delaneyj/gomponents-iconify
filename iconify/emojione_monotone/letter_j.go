package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterJ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m9.095 35.334c0 2.469-.417 4.424-1.251 5.867c-1.412 2.441-3.959 3.662-7.642 3.662s-6.153-.988-7.411-2.963c-1.256-1.975-1.885-4.67-1.885-8.082v-.668h5.729v.668c.049 2.244.279 3.822.693 4.734c.412.912 1.297 1.369 2.652 1.369c1.344 0 2.231-.498 2.663-1.496c.259-.592.388-1.588.388-2.992V17.139h6.063v20.195z"/>`),
		g.Group(children),
	)
}