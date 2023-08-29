package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M428 135q0 18-12.5 30.5T385 178q-17 0-29.5-12.5T343 135q0-17 12.5-29.5T385 93q18 0 30.5 12.5T428 135zm-17 72l-52 44q1 4 1 11q0 28-20 48t-48 20q-24 0-42.5-15T226 277L98 234q-17 12-36 12q-25 0-42.5-17.5T2 186t17.5-42.5T62 126q21 0 37 13.5t21 33.5l123 41q20-19 47-20l19-56v-3q0-32 22.5-54.5T386 58q31 0 53.5 22.5T462 135q0 25-14 44.5T411 207zM99 166q-12-22-36-22q-18 0-30 12.5T21 186t12 29.5T63 228q4 0 12-2l-20-7q-12-6-16-19t2-25q7-12 20-16t25 2zm227-31q0 25 17 42.5t42 17.5t42.5-17.5T445 135t-17.5-42T385 76t-42 17t-17 42zm1 133q0-19-13.5-32T281 223q-7 0-10 1l26 8q12 7 16 20t-3 25t-20 16t-25-3l-27-9q10 32 43 32q19 0 32.5-13t13.5-32z"/>`),
		g.Group(children),
	)
}