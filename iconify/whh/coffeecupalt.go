package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffeecupalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 256h-6l-26 320H573q-11-55-55-91.5T416 448t-102 36.5t-55 91.5H96L70 256h-6q-27 0-45.5-18.5T0 192.5T18.5 147T64 128V64q0-27 18.5-45.5T128 0h576q26 0 45 18.5T768 64v64q26 0 45 19t19 45.5t-19 45t-45 18.5zM259 640q11 55 55 91.5T416 768t102-36.5t55-91.5h158l-27 320q-4 24-21.5 44t-42.5 20H192q-25 0-42.5-20T128 960l-27-320h158z"/>`),
		g.Group(children),
	)
}