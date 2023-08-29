package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicScore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M491 0h-39q-21 7-34 24t-13 38v296q-29-17-64-17q-44 0-75 25.5T235 427t31 60t75 25q45 0 76-25t31-60q0-9-2-13q0-1 1-3.5t1-5.5V62q0-16 13-19h30q9 0 15-6t6-16q0-9-6-15t-15-6zM341 469q-26 0-45-12.5T277 427q0-18 19-30.5t45-12.5t45 12.5t19 30.5q0 17-19 29.5T341 469zM21 107h320q22 0 22-22q0-9-6-15t-16-6H21q-9 0-15 6T0 85q0 10 6 16t15 6zm0 85h320q10 0 16-6t6-15q0-22-22-22H21q-9 0-15 6t-6 16q0 9 6 15t15 6zm0 85h320q22 0 22-21t-22-21H21q-21 0-21 21t21 21zm0 86h192l43-43H21q-9 0-15 6t-6 15q0 10 6 16t15 6zm0 85h171v-43H21q-9 0-15 6t-6 16q0 9 6 15t15 6z"/>`),
		g.Group(children),
	)
}