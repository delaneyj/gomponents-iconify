package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M133 469q-17 0-29.5 13T91 512h170q0-17-12.5-30T219 469q-22-6-22-21v-45q30-4 49-15l-27-32q-18 7-43 7q-52 0-90-37.5T48 235v-64q0-7-2-11l-7-6q-7-5-12-5q-8 0-15 7t-7 15v64q0 65 43 112.5T155 403v45q0 7-5.5 12t-11.5 7zm43-128q17 0 26-4l-35-38q-24-4-39.5-21.5T112 237l-43-49v47q0 44 31 75t76 31zm107-104V107q0-45-31-76T176 0q-43 0-75 30l28 32q22-19 47-19q27 0 45.5 18t18.5 46v81zm42-88q-8 0-14.5 7t-6.5 15v64q0 15-2 23l34 39q11-37 11-62v-64q0-8-7-15t-15-7zm5 271q6 7 17 7q7 0 15-5q16-16 2-29l-49-56l-30-32l-15-17l-32-36l-126-145l-36-41l-32-38q-10-7-17-7q-8 0-15 5q-17 15-2 29l59 69l43 47l102 115l28 32l15 17l28 32z"/>`),
		g.Group(children),
	)
}