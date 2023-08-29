package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240.3 396.8c3.3 5.1 9.1 8.5 15.7 8.5s12.4-3.4 15.8-8.5L382 226.6c14.8-22.9 23.4-48.1 23.4-77.3C405.3 64.9 339 0 256 0S106.7 64.9 106.7 149.3c0 29.2 8.6 54.4 23.4 77.3l110.2 170.2zM256 64c47.1 0 85.3 38.2 85.3 85.3s-38.2 85.3-85.3 85.3s-85.3-38.2-85.3-85.3S208.9 64 256 64zm109.4 259.5L256 469.3L146.6 323.5c-37.4 19.6-61.3 48.9-61.3 81.8C85.3 464.2 161.7 512 256 512s170.7-47.8 170.7-106.7c0-32.9-23.9-62.2-61.3-81.8z"/>`),
		g.Group(children),
	)
}