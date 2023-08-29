package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M38 48h-6.1V25c-2.2 2.1-4.9 3.6-7.9 4.6v-5.5c1.6-.5 3.3-1.5 5.2-3c1.9-1.5 3.2-3.2 3.8-5.1h5v32"/>`),
		g.Group(children),
	)
}