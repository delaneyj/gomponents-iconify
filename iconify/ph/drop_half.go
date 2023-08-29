package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174 47.75a254.19 254.19 0 0 0-41.45-38.3a8 8 0 0 0-9.18 0A254.19 254.19 0 0 0 82 47.75C54.51 79.32 40 112.6 40 144a88 88 0 0 0 176 0c0-31.4-14.51-64.68-42-96.25ZM200 144a70.57 70.57 0 0 1-.46 8H136v-16h63.64q.36 4 .36 8Zm-16.61-56H136V72h36.89a175.85 175.85 0 0 1 10.5 16ZM136 200h37.19A71.67 71.67 0 0 1 136 215.54Zm0-16v-16h59.87a72 72 0 0 1-8 16Zm0-64v-16h55.39a116.84 116.84 0 0 1 5.45 16Zm23.89-64H136V32.6A257.22 257.22 0 0 1 159.89 56ZM56 144c0-50 42.26-92.71 64-111.4v182.94A72.08 72.08 0 0 1 56 144Z"/>`),
		g.Group(children),
	)
}