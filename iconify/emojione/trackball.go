package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trackball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d0d0" d="M52 62V18.9c0-8.4-4-15-14-15c-7.3 0-12 5.6-12 13.1h-2c0-9.4 5.7-15 14-15c10.7 0 16 6.6 16 16.9V62h-2"/><path fill="#3e4347" d="M40 33.1c0 16-6.7 28.9-15 28.9S10 49.1 10 33.1c0-9.8 6.7-22.7 15-22.7s15 12.9 15 22.7"/><path fill="#94989b" d="M15.2 25.5c-.7 2.5-1.2 5.1-1.2 7.6c0 4.9.6 9.3 1.6 13.1c.3-.5.6-1 1-1.6c5.8-10.5 2.5-16.1-1.4-19.1m19.6 0c-3.9 2.9-7.2 8.6-1.4 19c.3.6.7 1.1 1 1.6c1-3.8 1.6-8.3 1.6-13.1c0-2.4-.4-5-1.2-7.5"/><ellipse cx="25" cy="21.7" fill="#ed4c5c" rx="6" ry="5.6"/>`),
		g.Group(children),
	)
}