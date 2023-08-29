package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpsalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 576q-27 0-45.5-18.5T896 512q0-104-51.5-192.5t-140-140T512 128q-26 0-45-18.5t-19-45T467 19t45-19q104 0 199 40.5t163.5 109t109 163.5t40.5 199q0 27-18.5 45.5T960 576zm-191.5 0q-26.5 0-45.5-18.5T704 512q0-80-56-136t-136-56q-26 0-45-18.5t-19-45t19-45.5t45-19q87 0 160.5 43T789 351.5T832 512q0 27-18.5 45.5t-45 18.5zM640 512q0 66-54 104l227 228q19 18 19 44.5T813 934q-34 35-99 56t-121 27.5t-113 6.5q-55 0-112-10t-125-40.5T128 896T51 781T10.5 655.5T0 544q0-58 6-114t27.5-120.5T90 210q19-19 45.5-19t44.5 19l228 228q38-54 104-54q53 0 90.5 37.5T640 512z"/>`),
		g.Group(children),
	)
}