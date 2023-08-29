package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SweatDroplets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#75d6ff" d="M22.5 2.6c5.7 12.9 12.4 22.2 19.8 26.5c6.3 3.6 14.4 1.5 18-4.7c3.6-6.3 1.4-14.3-4.9-17.9c-7.5-4.3-19.1-5.4-32.9-3.9M12.2 11C5.7 19.8 2 27.9 2 34.6c0 5.7 4.6 10.3 10.2 10.3s10.2-4.6 10.2-10.3c0-6.7-3.9-14.9-10.2-23.6M29 31.2c-1.2 10.9-.3 19.8 3 25.6c2.8 4.9 9.1 6.6 14 3.8s6.5-9.1 3.7-14c-3.4-5.8-10.9-11-20.7-15.4"/>`),
		g.Group(children),
	)
}