package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moped(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 128a39.3 39.3 0 0 0-6.27.5l-34.24-91.31A8 8 0 0 0 168 32h-32a8 8 0 0 0 0 16h26.46l32.3 86.13a40.13 40.13 0 0 0-18 25.87h-40.22l-25-66.81A8 8 0 0 0 104 88H24a8 8 0 0 0 0 16h8v13.39A56.12 56.12 0 0 0 0 168a8 8 0 0 0 8 8h8.8a40 40 0 0 0 78.4 0h81.6a40 40 0 1 0 39.2-48Zm-173.33 2.27a8 8 0 0 0 5.33-7.54V104h50.46l21 56H16.81a40.07 40.07 0 0 1 25.86-29.73ZM56 192a24 24 0 0 1-22.62-16h45.24A24 24 0 0 1 56 192Zm160 0a24 24 0 0 1-15.43-42.36l7.94 21.17a8 8 0 0 0 15-5.62L215.55 144h.45a24 24 0 0 1 0 48Z"/>`),
		g.Group(children),
	)
}