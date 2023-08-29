package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PigNose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fc97b2" d="M62 37.3c0 17.3-13.4 15.3-30 15.3s-30 2-30-15.3s13.4-25.9 30-25.9S62 20 62 37.3"/><path fill="#a5627a" d="M25.9 33.6c0 7.4-3.3 13.4-7.3 13.4c-4.1 0-7.3-6-7.3-13.4s3.3-13.4 7.3-13.4c4.1.1 7.3 6 7.3 13.4"/><ellipse cx="19.9" cy="33.6" fill="#bf7b90" rx="6.1" ry="12.8"/><path fill="#a5627a" d="M52.7 33.6c0 7.4-3.3 13.4-7.3 13.4c-4.1 0-7.3-6-7.3-13.4s3.3-13.4 7.3-13.4c4 .1 7.3 6 7.3 13.4"/><ellipse cx="46.7" cy="33.6" fill="#bf7b90" rx="6.1" ry="12.8"/>`),
		g.Group(children),
	)
}