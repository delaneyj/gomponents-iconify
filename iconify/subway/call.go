package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Call(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m426.7 453.8l-38.1-79.1c-8.2-16.9-18.8-29.2-37.1-21.7l-36.1 13.4c-28.9 13.4-43.3 0-57.8-20.2l-65-147.9c-8.2-16.9-3.9-32.8 14.4-40.3l50.5-20.2c18.3-7.6 15.4-23.4 7.2-40.3l-43.3-80.6c-8.2-16.9-25-21-43.3-13.5c-36.6 15.1-66.9 38.8-86.6 73.9c-24 42.9-12 102.6-7.2 127.7c4.8 25.1 21.6 69.1 43.3 114.2c21.7 45.2 40.7 80.7 57.8 100.8c17 20.1 57.8 75.1 108.3 87.4c41.4 10 86.1 1.6 122.7-13.5c18.4-7.2 18.4-23.1 10.3-40.1z"/>`),
		g.Group(children),
	)
}