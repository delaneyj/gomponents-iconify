package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 128a52 52 0 1 0-88 37.47V224a4 4 0 0 0 5.79 3.58L192 212.47l30.21 15.11a4.05 4.05 0 0 0 1.79.42a4 4 0 0 0 4-4v-58.53A51.85 51.85 0 0 0 244 128Zm-96 0a44 44 0 1 1 44 44a44.05 44.05 0 0 1-44-44Zm72 89.53l-26.21-13.11a4 4 0 0 0-3.58 0L164 217.53v-45.75a51.8 51.8 0 0 0 56 0ZM132 192a4 4 0 0 1-4 4H40a12 12 0 0 1-12-12V56a12 12 0 0 1 12-12h176a12 12 0 0 1 12 12a4 4 0 0 1-8 0a4 4 0 0 0-4-4H40a4 4 0 0 0-4 4v128a4 4 0 0 0 4 4h88a4 4 0 0 1 4 4Zm-16-56a4 4 0 0 1-4 4H72a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Zm0-32a4 4 0 0 1-4 4H72a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}