package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jcore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1016 64L899 281q4 17 11 27t13 11l6 1q13 0 22.5 9.5T961 352t-9.5 22.5T929 384q-2 0-7 .5t-19.5 4.5t-28.5 10.5t-33 21t-35 34.5l-65 121q-14 27-43 45.5T643 640H39q-26 0-35-18.5t6-45.5l117-218q-4-16-11-26t-13-11l-6-1q-13 0-22.5-9.5T65 288t9.5-22.5T97 256q3 0 7.5-.5t18.5-4t28-10t33-21t35-34.5l66-122q14-27 43-45.5T383 0h604q26 0 35 18.5t-6 45.5zM385 448q-11 23-22.5 37t-25 19.5T315 511t-26 1v64h66q44 0 88-37.5t74-90.5l139-256H526zM595 64l-34 64h130l34-64H595z"/>`),
		g.Group(children),
	)
}