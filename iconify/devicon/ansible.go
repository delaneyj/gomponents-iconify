package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ansible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#1A1918" d="M126 64c0 34.2-27.8 62-62 62S2 98.2 2 64S29.8 2 64 2s62 27.8 62 62"/><path fill="#FFF" d="m65 39.9l16 39.6l-24.1-19.1L65 39.9zm28.5 48.7L68.9 29.2c-.7-1.7-2.1-2.6-3.8-2.6c-1.7 0-3.2.9-3.9 2.6L34 94.3h9.3L54 67.5l32 25.9c1.3 1 2.2 1.5 3.4 1.5c2.4 0 4.5-1.8 4.5-4.4c.1-.5-.1-1.2-.4-1.9z"/>`),
		g.Group(children),
	)
}