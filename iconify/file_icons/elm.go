package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 241.633l110.367-110.368H144.98L256 241.633zM13.714 0l111.674 111.02h242.285L256 0H13.714zm377.47 134.53l120.163 120.817L390.53 376.163L269.714 256l121.47-121.47zM512 227.266V0H284.082L512 227.265zM241.633 256L0 13.714v483.919L241.633 256zm162.612 134.53L512 497.634V282.776L404.245 390.53zM256 269.715L13.714 512h483.919L256 269.714z"/>`),
		g.Group(children),
	)
}