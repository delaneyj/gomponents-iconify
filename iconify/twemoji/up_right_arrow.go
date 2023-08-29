package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M0 32a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4a4 4 0 0 0-4 4v28z"/><path fill="#FFF" d="M27 25V9H11z"/><path fill="#FFF" d="M7 23.343L19.816 10.53l5.656 5.657L12.657 29z"/>`),
		g.Group(children),
	)
}