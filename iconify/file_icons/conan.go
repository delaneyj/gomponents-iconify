package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M222.693 207.677c-46.788-2.9-151.013-62.967-49.6-121.072c83.426-47.799 160.802-14.282 208.409 36.792c0 0-27.37 16.852-59.206 32.04c22.495-50.286-61.3-78.374-109.355-52.24c-48.054 26.134-52.773 79.742 65.474 76.619zM480.518 354.38V119.357L235.38 0L0 113.299V340.01L242.727 512zM235.276 11.979L465.77 127.766L235.623 258.997L12.536 121.051zm230.491 129.883v202.594l-91.863 59.74V195.59zm-106.963 63.481V414.61l-116.11 73.25V271.34z"/>`),
		g.Group(children),
	)
}