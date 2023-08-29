package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a5.5 5.5 0 0 0-5.49 5.15l-.048.783l-.77.14A4.502 4.502 0 0 0 6.5 19h11a4.5 4.5 0 0 0 .809-8.928l-.771-.14l-.049-.781A5.5 5.5 0 0 0 12 4ZM4.598 8.283a7.502 7.502 0 0 1 14.804 0A6.502 6.502 0 0 1 17.5 21h-11A6.5 6.5 0 0 1 4.598 8.283ZM12 7.086l4.414 4.414L15 12.914l-2-2V17h-2v-6.086l-2 2L7.586 11.5L12 7.086Z"/>`),
		g.Group(children),
	)
}