package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addfriend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 896h-64v64q0 27-19 45.5t-45.5 18.5t-45-18.5T768 960v-64h-64q-27 0-45.5-19T640 831.5t18.5-45T704 768h64v-64q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v64h64q26 0 45 18.5t19 45t-19 45.5t-45 19zM710 640q-32 26-70 42v28q-28 10-46 34.5T576 800v64q0 40 28 68t68 28h32v62q-70 2-192 2h-76l-81-1l-87-2.5l-80-4.5l-75.5-8.5l-57.5-12L13.5 979L0 957q2-88 110-155.5T384 712v-33q-52-23-90-65t-60-98.5t-32-121T192 256q0-65 25-114.5t69-80t101-46T512 0t125 15.5t101 46t69 80T832 256q0 220-75 330q-34 18-47 54z"/>`),
		g.Group(children),
	)
}