package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSunFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 72a76.2 76.2 0 0 0-20.26 2.73a55.63 55.63 0 0 0-9.41-11.54l9.51-13.57a8 8 0 1 0-13.11-9.18L113.22 54A55.9 55.9 0 0 0 88 48h-1.74l-2.89-16.29a8 8 0 1 0-15.75 2.77l2.88 16.34a56.1 56.1 0 0 0-23.27 14.85l-13.62-9.53a8 8 0 1 0-9.17 13.11L38 78.77A55.55 55.55 0 0 0 32 104v1.72l-16.29 2.88a8 8 0 0 0 1.38 15.88a8.24 8.24 0 0 0 1.39-.12l16.32-2.88a55.74 55.74 0 0 0 5.86 12.42A52 52 0 0 0 76 224h80a76 76 0 0 0 0-152Zm-71.08 48.76a52.14 52.14 0 0 0-31 4.17a40 40 0 0 1 66.62-44.17a76.26 76.26 0 0 0-35.62 40Z"/>`),
		g.Group(children),
	)
}