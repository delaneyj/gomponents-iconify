package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.627 1.087l.558-.004l6.091 4.037l-.003.836L8.182 9.92l-.551-.004l-5.91-3.963l-.003-.828l5.91-4.037Zm.286 1.016l-5.021 3.43l5.021 3.368l5.176-3.368l-5.176-3.43ZM1.793 8.5l5.838 3.915l.55.004L14.206 8.5h-1.833l-4.459 2.9l-4.325-2.9H1.793Zm5.838 6.415L1.793 11h1.795l4.325 2.9l4.459-2.9h1.833l-6.023 3.92l-.551-.005Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}