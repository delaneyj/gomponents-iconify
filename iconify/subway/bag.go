package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M421.2 128h-42.7v21.3c0 23.5-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7V128h-85.3v21.3c0 23.5-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7V128H79.9c0 213.3-21.3 384-21.3 384h384c-.1 0-21.4-170.7-21.4-384zm-256 42.7c11.8 0 21.3-9.5 21.3-21.3v-42.7c0-35.4 28.6-64 64-64s64 28.6 64 64v42.7c0 11.8 9.5 21.3 21.3 21.3s21.3-9.5 21.3-21.3v-42.7C357.2 47.8 309.4 0 250.5 0c-58.9 0-106.7 47.8-106.7 106.7v42.7c.1 11.7 9.6 21.3 21.4 21.3z"/>`),
		g.Group(children),
	)
}