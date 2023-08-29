package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 86q24-4 52-4q36 0 54 15q19 14 19 41q0 26-16 42q-22 20-59 20q-7 0-17-2v70H2V86zm33 86q5 1 17 1q41 0 41-34q0-31-38-31q-12 0-20 2v62zm103 96V135h34v133h-34zm57-94q0-8-.5-21t-.5-18h29l2 20h1q13-23 43-23q20 0 34 14t14 43v79h-34v-75q0-34-26-34q-19 0-27 20q-1 3-1 11v78h-34v-94zm266 75q0 41-20 60q-17 16-52 16q-28 0-47-11l8-25q16 10 39 10q39 0 39-40v-12h-1q-12 20-39 20q-25 0-41-18t-16-47q0-32 17.5-51t43.5-19t39 21l2-18h29q0 3-.5 15.5T461 173v76zm-34-61q0-6-1-9q-5-21-28-21q-15 0-24 11.5t-9 31.5q0 18 9 29.5t24 11.5q22 0 28-21q1-3 1-11v-22zM155 59q-13 0-22 8.5T124 89t9 21.5t22 8.5t22-8.5t9-21.5t-9-21.5t-22-8.5zm0 54q-10 0-17-7t-7-17t7-17t17-7t17 7t7 17t-7 17t-17 7zm11-16q0 5-4 5q-5 2-5-4q0-5 4-5h2V81l-12 2v16q0 5-4 5q-5 0-5-4q0-5 4-5h2V77l18-3v23z"/>`),
		g.Group(children),
	)
}