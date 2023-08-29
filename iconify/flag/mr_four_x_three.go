package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#cd2a3e" d="M0 0h640v480H0z"/><path fill="#006233" d="M0 72h640v336H0z"/><path fill="#ffc400" d="M470 154.6a150 150 0 0 1-300 0a154.9 154.9 0 0 0-5 39.2a155 155 0 1 0 310 0a154.4 154.4 0 0 0-5-39.2z" class="mr-st1"/><path fill="#ffc400" d="m320 93.8l-13.5 41.5H263l35.3 25.6l-13.5 41.4l35.3-25.6l35.3 25.6l-13.5-41.4l35.3-25.6h-43.6z"/>`),
		g.Group(children),
	)
}