package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneTiltThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224.47 31.52a11.87 11.87 0 0 0-11.82-3L20.74 86.67a12 12 0 0 0-1.88 22.4l85.61 40.55a4 4 0 0 1 1.91 1.9l40.55 85.62a11.88 11.88 0 0 0 10.81 6.86h1a11.89 11.89 0 0 0 10.52-8.63l58.21-192v-.08a11.85 11.85 0 0 0-3-11.77Zm-4.62 9.54l-58.23 192a4 4 0 0 1-7.46.63l-40.55-85.59l49.24-49.23a4 4 0 1 0-5.66-5.66l-49.24 49.24l-85.66-40.61a3.88 3.88 0 0 1-2.27-4a3.93 3.93 0 0 1 3-3.54L214.9 36.16A3.93 3.93 0 0 1 216 36a4 4 0 0 1 2.79 1.19a3.93 3.93 0 0 1 1.06 3.87Z"/>`),
		g.Group(children),
	)
}