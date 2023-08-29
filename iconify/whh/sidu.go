package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sidu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M596 320q79 53 125.5 137.5T768 640q0 104-51.5 192.5t-140 140T384 1024t-192.5-51.5t-140-140T0 640q0-99 46.5-183T172 320Q93 267 46.5 182.5T0 0q124 0 223 71.5T362 257q19-1 22-1t22 1q40-114 139-185.5T768 0q0 98-46.5 182.5T596 320zm-212 64q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75z"/>`),
		g.Group(children),
	)
}