package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246 128a54 54 0 1 0-92 38.32V224a6 6 0 0 0 8.68 5.37L192 214.71l29.32 14.66a6 6 0 0 0 2.68.63a5.93 5.93 0 0 0 3.15-.9A6 6 0 0 0 230 224v-57.68A53.83 53.83 0 0 0 246 128Zm-96 0a42 42 0 1 1 42 42a42 42 0 0 1-42-42Zm68 86.29l-23.32-11.66a6 6 0 0 0-5.36 0L166 214.29v-39a53.87 53.87 0 0 0 52 0ZM134 192a6 6 0 0 1-6 6H40a14 14 0 0 1-14-14V56a14 14 0 0 1 14-14h176a14 14 0 0 1 14 14a6 6 0 0 1-12 0a2 2 0 0 0-2-2H40a2 2 0 0 0-2 2v128a2 2 0 0 0 2 2h88a6 6 0 0 1 6 6Zm-16-56a6 6 0 0 1-6 6H72a6 6 0 0 1 0-12h40a6 6 0 0 1 6 6Zm0-32a6 6 0 0 1-6 6H72a6 6 0 0 1 0-12h40a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}