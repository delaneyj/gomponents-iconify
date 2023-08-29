package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469.3 95.3H448c-23.5 0-42.7 19.1-42.7 42.7v52l-152.7-89.8c-21.6-12.7-39.2.1-39.2 28.5v73.9L39.2 100.1C17.6 87.4 0 100.2 0 128.6v274.8c0 28.3 17.6 41.2 39.2 28.5l174.2-102.5v74c0 28.3 17.6 41.2 39.2 28.5l152.7-89.8V394c0 23.5 19.1 42.7 42.7 42.7h21.3c23.5 0 42.7-19.1 42.7-42.7V138c0-23.6-19.1-42.7-42.7-42.7z"/>`),
		g.Group(children),
	)
}