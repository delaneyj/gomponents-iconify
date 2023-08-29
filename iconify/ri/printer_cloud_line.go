package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterCloudLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2a1 1 0 0 1 1 1v4h3a1 1 0 0 1 1 1v5.324a5.925 5.925 0 0 0-.61-.713A5.595 5.595 0 0 0 20 11.583V9H4v8h2v-1a1 1 0 0 1 1-1h5.194a5.305 5.305 0 0 0-.07.283A4.945 4.945 0 0 0 10.562 17H8.001v3h2.054a4.48 4.48 0 0 0 .817 2H7a1 1 0 0 1-1-1v-2H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h3V3a1 1 0 0 1 1-1h10Zm-1 2H8v3h8V4Zm-8 6v2H5v-2h3Zm13 6.5a3.5 3.5 0 1 0-7 0l.003.102a2.751 2.751 0 0 0 .58 5.393l.167.005h5.5l.168-.005a2.75 2.75 0 0 0 .58-5.392L21 16.5Zm-4.993-.145a1.5 1.5 0 0 1 2.986 0L19 16.5v1.62c.696.199 1.177.334 1.444.406A.75.75 0 0 1 20.255 20h-5.51a.75.75 0 0 1-.19-1.474c.238-.064.645-.178 1.22-.342L16 18.12V16.5l.007-.145Z"/>`),
		g.Group(children),
	)
}