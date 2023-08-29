package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 255v-63q0-10 10-10h111l-49 83H12q-10 0-10-10zm78 128l-35 24q-3 1-4.5 1t-2-1.5t-.5-4.5l4-41q2-10 11-5l27 16q9 5 0 11zm9-23l-28-16q-8-5-3-14l126-219q5-9 14-4l28 16q8 5 3 14L103 356q-5 9-14 4zm96-95l48-83h28l39 83H185zm157 25q-13-6-80-163q-2-5-7.5-17T245 88.5t-9.5-23t-9-22.5t-6.5-19.5t-2.5-15T220 1q4-2 15 13.5T264 61t27 44q11 18 31 54.5t37 69t19 35.5q4 8-1.5 15T364 289q-12 7-22 1zm35 58l-16-26q-5-10 4-15l16-8q10-6 15 5l13 25q6 10-5 16l-11 6q-8 7-16-3zm61 67q-2-5-19-12t-23-16q-12-17 3-29q21-12 30.5 2.5T438 395v20zm24-160q0 10-10 10h-55q-2-6-4-9q-1-3-10-19q-4-8-14.5-28T354 182h98q10 0 10 10v63z"/>`),
		g.Group(children),
	)
}