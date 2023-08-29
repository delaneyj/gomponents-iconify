package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFoldedFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201 48h-21.65l-13.69-13.66A8 8 0 0 0 160 32H96a8 8 0 0 0-5.66 2.34L76.65 48H55a15 15 0 0 0-15 15v154a15 15 0 0 0 15 15h61a4 4 0 0 0 4-4V112.27a8.18 8.18 0 0 1 7.47-8.25a8 8 0 0 1 8.53 8V228a4 4 0 0 0 4 4h61a15 15 0 0 0 15-15V63a15 15 0 0 0-15-15ZM86.54 115.08A4 4 0 0 1 80 112V67.31l15.24-15.24l23.47 35.21ZM128 88Zm48 24a4 4 0 0 1-2.3 3.63a3.93 3.93 0 0 1-4.21-.51l-32.2-27.82l23.47-35.21L176 67.31Z"/>`),
		g.Group(children),
	)
}