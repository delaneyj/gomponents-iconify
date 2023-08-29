package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeakOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 0q0 30-11 57l-34-34q3-11 3-23h42zM0 27L27 0l357 357l-27 27l-61-61q-19 28-19 61h-42q0-51 31-91l-31-30q-43 52-43 121h-43q0-86 56-152l-53-53Q86 235 0 235v-43q68 0 122-43l-31-31q-40 31-91 31v-42q33 0 61-19zM235 0q0 64-34 120l-31-31q22-42 22-89h43zm126 280l-34-34q28-11 57-11v42q-12 0-23 3zm-97-97q56-34 120-34v43q-47 0-89 22z"/>`),
		g.Group(children),
	)
}