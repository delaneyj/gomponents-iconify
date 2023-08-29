package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M251 85.3c-117.8 0-213.3 95.5-213.3 213.3C37.7 416.5 133.2 512 251 512c117.8 0 213.3-95.5 213.3-213.3c0-117.8-95.5-213.4-213.3-213.4zM357.7 320h-128V149.3h42.7v128h85.3V320zm100-171.9c17.1-15.6 27.9-37.8 27.9-62.8c0-47.1-38.2-85.3-85.3-85.3c-35.6 0-66 21.8-78.8 52.7c55.5 15.9 103.1 50 136.2 95.4zM180.5 52.7C167.6 21.8 137.2 0 101.7 0C54.5 0 16.3 38.2 16.3 85.3c0 24.9 10.9 47.2 27.9 62.8c33.2-45.4 80.8-79.5 136.3-95.4z"/>`),
		g.Group(children),
	)
}