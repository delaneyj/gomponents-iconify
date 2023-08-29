package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 128a52 52 0 1 0-92 33.19V220a12 12 0 0 0 18 10.42l22-12.6l22.05 12.6A12 12 0 0 0 236 220v-58.81A51.77 51.77 0 0 0 248 128Zm-52-28a28 28 0 1 1-28 28a28 28 0 0 1 28-28Zm6 93.58a12 12 0 0 0-11.9 0l-10.1 5.74v-21.85a51.86 51.86 0 0 0 32 0v21.85ZM140 192a12 12 0 0 1-12 12H40a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h176a20 20 0 0 1 20 20a12 12 0 0 1-23.32 4H44v120h84a12 12 0 0 1 12 12Zm-12-52a12 12 0 0 1-12 12H76a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12Zm0-40a12 12 0 0 1-12 12H76a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}