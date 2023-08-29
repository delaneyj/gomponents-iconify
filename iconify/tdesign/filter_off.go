package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1.586L21.414 20L20 21.414l-5.5-5.5V21h-5v-8.182L4.933 6.347L1.586 3L3 1.586Zm8.5 11.328V19h1v-5.086l-1-1ZM6.545 3H21.43l-6.034 8.549l-1.634-1.153L17.57 5H9.51l-.899.018L6.545 3Z"/>`),
		g.Group(children),
	)
}