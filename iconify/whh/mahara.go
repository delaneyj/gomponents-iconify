package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mahara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 768q0 70-43 128.5t-116.5 93T320 1024t-160.5-34.5t-116.5-93T0 768q0-81 70.5-188.5T226 418q-43-26-70.5-67T128 256q0-35 25.5-81.5t58.5-84t64.5-64T320 0t43.5 26.5t64.5 64t58.5 84T512 256q0 54-27.5 95T414 418q85 54 155.5 161.5T640 768zM256 512q-3 2-7 6t-15 17t-19.5 27t-15.5 35t-7 43q0 31 27.5 47.5T288 704q17 0 24.5-12.5T320 640q0-29-14-60.5T277 530t-21-18zm80-256q-9 0-12.5 8.5T320 299q0 19 7 40t14.5 33t10.5 12q4-4 9-11t14-29.5t9-44.5q0-21-13.5-32T336 256z"/>`),
		g.Group(children),
	)
}