package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyPostOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M19.758 32.97a1 1 0 0 1-.728-1.212L19.72 29H17a1 1 0 1 1 0-2h3.22l.81-3.242a1 1 0 1 1 1.94.485L22.28 27H25a1 1 0 1 1 0 2h-3.22l-.81 3.243a1 1 0 0 1-1.212.727Z"/><path fill-rule="evenodd" d="M9.014 20.138C9.402 18.869 10.59 18 11.938 18H36v-2h-7v-6h9v8.962l5.935 13.687a.995.995 0 0 1-.642 1.303l.001.004L33.5 37l1-2.5l7.092-2.122l-4.607-10.25l-4.803 13.734C31.795 37.131 30.607 38 29.26 38H7.054c-2.044 0-3.51-1.937-2.923-3.862l4.883-14ZM6.036 34.749l4.88-13.99l.011-.037A1.047 1.047 0 0 1 11.938 20h23.673l-5.33 15.24l-.011.038a1.047 1.047 0 0 1-1.011.722H7.054c-.737 0-1.178-.663-1.018-1.25ZM36 12h-5v2h5v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}