package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dango(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M474.2 37.8c-21.6-21.6-56.2-21.6-77.8 0c-21.6 21.6-21.6 56.2 0 77.8c21.6 21.6 56.2 21.6 77.8 0c21.6-21.6 21.6-56.2 0-77.8zm-90.5 90.5a54.984 54.984 0 0 0-77.8 0c-21.6 21.6-21.6 56.2 0 77.8c21.6 21.6 56.2 21.6 77.8 0c21.5-21.6 21.5-56.2 0-77.8zm-90.5 90.5a54.984 54.984 0 0 0-77.8 0c-21.6 21.6-21.6 56.2 0 77.8c21.6 21.6 56.2 21.6 77.8 0c21.5-21.6 21.5-56.2 0-77.8zm-96.5 83.7L21.62 477.6l12.73 12.8L209.5 315.3c-2.5-1.9-4.7-3.8-6.8-6c-2.2-2.1-4.1-4.3-6-6.8z"/>`),
		g.Group(children),
	)
}