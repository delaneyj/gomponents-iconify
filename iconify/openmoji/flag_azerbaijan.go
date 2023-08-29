package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAzerbaijan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 42h62v13H5z"/><path fill="#61b2e4" d="M5 17h62v13H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m38.436 34.61l.586-1.61l.562 1.618l1.553-.724l-.747 1.542l1.61.586l-1.618.562l.724 1.553l-1.542-.747l-.586 1.61l-.562-1.618l-1.553.724l.747-1.542l-1.61-.586l1.618-.562l-.724-1.553l1.542.747zM30.53 36a4.053 4.053 0 0 1 3.38-3.924a4.396 4.396 0 0 0-.812-.076a4.004 4.004 0 1 0 0 8a4.396 4.396 0 0 0 .811-.076A4.053 4.053 0 0 1 30.531 36Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}