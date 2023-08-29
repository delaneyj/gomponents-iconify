package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaClock0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaClock1" stroke-opacity="1" d="M 80.039519,211.06969 64.056987,239.90745"/><path id="galaClock2" stroke-opacity="1" d="m 175.96058,211.06969 15.98254,28.83776"/><circle id="galaClock3" cx="128" cy="128.007" r="95.915"/><path id="galaClock4" d="M 35.294352,102.43866 A 47.957299,47.957299 0 0 1 17.212686,53.792007 47.957299,47.957299 0 0 1 53.990946,17.175027 47.957299,47.957299 0 0 1 102.55767,35.470309"/><path id="galaClock5" stroke-opacity="1" d="m 127.99967,32.092482 3.8e-4,-15.985761"/><path id="galaClock6" stroke-opacity="1" d="M 128.00005,80.049788 V 128.00708"/><path id="galaClock7" stroke-opacity="1" d="m 128.00005,128.00708 33.91093,33.91093"/><path id="galaClock8" d="M 220.70575,102.43866 A 47.957299,47.957299 0 0 0 238.78742,53.792007 47.957299,47.957299 0 0 0 202.00916,17.175027 47.957299,47.957299 0 0 0 153.44244,35.470309"/></g>`),
		g.Group(children),
	)
}