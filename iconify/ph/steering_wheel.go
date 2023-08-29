package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteeringWheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 152a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm104-24A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-192 0v.33a135.93 135.93 0 0 1 176 0V128a88 88 0 0 0-176 0Zm67.5 85.58L90.45 168H49.63a88.35 88.35 0 0 0 57.87 45.58ZM128 216h2.49l20.07-53.57a16.07 16.07 0 0 1 15-10.39h47.12c.38-1.31.72-2.64 1-4a120 120 0 0 0-171.4 0c.31 1.34.65 2.67 1 4h47.17a16.08 16.08 0 0 1 15 10.4l20 53.56H128Zm78.37-48h-40.82l-17.09 45.59A88.34 88.34 0 0 0 206.37 168Z"/>`),
		g.Group(children),
	)
}