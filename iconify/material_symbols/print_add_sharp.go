package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintAddSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-4H2V8h20v3.75q-.675-.35-1.413-.55t-1.512-.2q-1.95 0-3.538 1.1T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H6ZM6 7V3h12v4H6Zm12 14v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}