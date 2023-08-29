package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 272.748H357.21c-30.383-117.105-176.84-106.081-202.42 0H0c13-296.638 475-323.304 512 0zm-452.465 35.72h95.256c26.876 104.976 168.209 116.309 202.418 0h95.256c-48.132 204.976-316.132 230.309-392.93 0zm261.953-17.86c0-50.225-54.746-81.787-98.306-56.674s-43.56 88.237 0 113.35s98.306-6.449 98.306-56.675z"/>`),
		g.Group(children),
	)
}