package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M62 23q16-13 43-18t43-2q27 4 29 26l1 160q0 9-4.5 16.5T163 214q-15 4-26-16L53 61q-1-2-4.5-6t-4-6t0-6t5-9T62 23zM34 326l85-29q21-10 23-26q2-18-19-27l-89-35q-9-4-17.5 3T5 233q-2 62 0 75q2 9 11 15.5t18 2.5zm147 15q1-9-3.5-16t-10.5-8q-14-2-30 16l-59 71q-6 7-3.5 20T85 440l62 22q9 3 20.5-4.5T180 441zm147-16l-74-21q-21-7-30 0q-10 8 2 27l47 86q5 8 18.5 7t18.5-9q28-41 33-57q4-11-1-21t-14-12zm18-120q-3-8-16.5-27T307 152q-21-18-35 0l-48 56q-15 18-4 34q9 16 32 11l82-15q20-5 12-33z"/>`),
		g.Group(children),
	)
}