package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TranslateThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m235.58 214.21l-56-112a4 4 0 0 0-7.16 0L149.55 148A92.05 92.05 0 0 1 94 127.36A99.68 99.68 0 0 0 123.91 60H152a4 4 0 0 0 0-8H92V32a4 4 0 0 0-8 0v20H24a4 4 0 0 0 0 8h91.91A91.8 91.8 0 0 1 88 122.05a92 92 0 0 1-22.77-35.38a4 4 0 1 0-7.54 2.66a99.59 99.59 0 0 0 24.3 38A91.59 91.59 0 0 1 24 148a4 4 0 0 0 0 8a99.54 99.54 0 0 0 64-23.21a100.09 100.09 0 0 0 57.66 23l-29.22 58.43a4 4 0 0 0 7.16 3.58l14.87-29.8h75.06l14.89 29.79A4 4 0 0 0 232 220a4.12 4.12 0 0 0 1.79-.42a4 4 0 0 0 1.79-5.37ZM142.47 180L176 112.94L209.53 180Z"/>`),
		g.Group(children),
	)
}