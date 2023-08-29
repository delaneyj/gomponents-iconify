package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 9V3a2 2 0 0 0-2-2c-1.12 0-2 .94-2 2v6c0 1.1.9 2 2 2c1.13 0 2-.94 2-2zm4 0c0 2.97-2.16 5.43-5 5.91V17h2c.56 0 1 .45 1 1s-.44 1-1 1H7c-.55 0-1-.45-1-1s.45-1 1-1h2v-2.09C6.17 14.43 4 11.97 4 9c0-.55.45-1 1-1c.56 0 1 .45 1 1a3.999 3.999 0 1 0 8 0c0-.55.45-1 1-1c.56 0 1 .45 1 1z"/>`),
		g.Group(children),
	)
}