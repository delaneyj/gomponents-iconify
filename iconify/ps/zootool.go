package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zootool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M23 329q22 0 47-23q53 101 77 128.5t60 27.5q39 0 72-24t33-51q0-19-13-26q-33 18-59 18q-18 0-37-20.5T132 260q81-72 129-142q50-33 50-65q0-14-9.5-26.5T279 14q-11 0-25 23Q167 2 121 2Q97 2 79.5 20.5T62 64q0 27 20 34q19-13 36-13q27 0 96 22q-15 22-49 58t-63 59q-20-17-37-17q-29 0-45 26.5T4 288q0 41 19 41z"/>`),
		g.Group(children),
	)
}