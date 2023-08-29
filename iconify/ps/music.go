package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M486 13Q463-5 431 2L218 64q-21 4-34 21.5T171 126v235q-26-20-64-20q-45 0-76 25.5T0 427t31 60t76 25q44 0 75-27t31-63v-1q-1-1-2-16h2V213q0-13 15-21l214-62q8-3 19 4q8 7 8 17v143q-29-17-64-17q-44 0-75 25.5T299 363t31 60t75 25q45 0 76-25t31-60q0-8-4-22h4V64q0-31-26-51zM107 469q-26 0-45-12.5T43 427q0-18 19-30.5t45-12.5t45 12.5t19 30.5q0 17-19 29.5T107 469zm298-64q-26 0-45-12.5T341 363q0-18 19-30.5t45-12.5t45 12.5t19 30.5q0 17-19 29.5T405 405zm26-313l-213 62q-3 0-5 2v-30q0-14 15-21l214-62q15 0 19 4q8 5 8 17v28q-18-8-38 0z"/>`),
		g.Group(children),
	)
}