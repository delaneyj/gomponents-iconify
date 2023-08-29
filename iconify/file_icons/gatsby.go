package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gatsby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M127.711 34.453c-170.281 98.17-170.281 344.925 0 443.094S512 452.338 512 256S297.992-63.716 127.711 34.453zM54.857 259.657l199.314 197.486c-105.53-2.169-196.602-82.508-199.314-197.486zm245.029 192L60.343 212.114c31.06-150.851 240.92-226.014 358.4-74.971l-27.429 23.771c-81.68-107.279-241.307-86.442-290.743 40.229l210.286 210.286c53.029-18.286 93.257-64 106.057-118.858h-87.771V256h128c0 95.086-67.657 175.543-157.257 195.657z"/>`),
		g.Group(children),
	)
}