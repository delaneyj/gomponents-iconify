package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 0q61 0 117 32.5t96.5 86.5t74 121.5t54 139.5T757 518.5T768 640q0 104-51.5 192.5t-140 140T384 1024t-192.5-51.5t-140-140T0 640q0-55 11-121.5T42.5 380t54-139.5t74-121.5T267 32.5T384 0z"/>`),
		g.Group(children),
	)
}