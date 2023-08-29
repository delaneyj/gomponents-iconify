package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M290.7 57.4L57.4 290.7c-25 25-25 65.5 0 90.5l80 80c12 12 28.3 18.7 45.3 18.7H512c17.7 0 32-14.3 32-32s-14.3-32-32-32H387.9l130.7-130.6c25-25 25-65.5 0-90.5L381.3 57.4c-25-25-65.5-25-90.5 0zm6.7 358.6H182.6l-80-80l124.7-124.7l137.4 137.4l-67.3 67.3z"/>`),
		g.Group(children),
	)
}