package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharethree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 512h-86q-40 113-139 184.5T512 768t-223-71.5T150 512h141q33 58 92 93t129 35q106 0 181-75t75-181h192q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zM640 384q0 53-37.5 90.5T512 512t-90.5-37.5T384 384t37.5-90.5T512 256t90.5 37.5T640 384zM512 128q-106 0-181 75t-75 181H64q-26 0-45-18.5t-19-45T19 275t45-19h86q40-113 139-184.5T512 0t223 71.5T874 256H734q-34-58-93-93t-129-35z"/>`),
		g.Group(children),
	)
}