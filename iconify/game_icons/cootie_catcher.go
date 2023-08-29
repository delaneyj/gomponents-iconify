package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CootieCatcher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M167 41.97L253.6 294l96.1-249.25l-77.9 72.15l-17.3 141.2l5.1-141.7zm199.1 4.82L278.3 271.5l78.2-116.4l85.7 57.8l-69.7-141.29l88.3 152.69l8.4-12.3zm-211.2 5.6C114 110.2 74.44 168.6 54.56 235l74.34-94.8l98 127.8zM127.7 167.2l-69.89 88.1c-27.16 94.9 4.45 210.4 4.45 210.4l176.34-88.3L128 214c-46 97.5-61.29 213-61.29 213s-2.74-158.3 59.79-245l111.9 173.1l4.7-38.6zm237.4 18.6L259.6 320.2l-5.8 35.6l111-153.6c75 121.2 75.3 241.6 75.3 241.6S406.1 310 361.4 230L253 380l188 90c41.5-104.6 26.4-216.6 26.4-216.6z"/>`),
		g.Group(children),
	)
}