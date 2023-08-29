package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M307.5 64.4c31.9 0 69.2 14 95.8 31.9V21.8C376.3 8.9 339.5.5 307.5.5C254.3.5 216.6 14.7 181 45.3c-35.6 30.6-63.1 71.4-75.8 125.5H62.7v63.9h42.6v42.6H62.7v63.9h42.6c21.3 117.1 95.8 170.3 202.3 170.3c42.6 0 62.8-4.2 95.8-21.3v-74.5c-32.6 20-65.9 31.9-95.8 31.9c-64.2 0-103-34-119.5-106.3l172.7-.1v-63.9H190.4v-42.6h170.3v-63.9H190.4c8.9-33.8 23.8-59.3 46.1-79.4c22.4-20.1 39.1-27 71-27z"/>`),
		g.Group(children),
	)
}