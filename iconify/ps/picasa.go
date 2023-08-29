package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M138 333h301q-26 55-76 89.5T253 462h-42q-40-3-73-19V333zM327 22Q281 2 232 2q-41 0-80 15q3 3 87.5 79.5T327 176V22zm-200 5q-58 30-91.5 85T2 232q0 28 8 60q3-2 102.5-92.5T214 107q-2-2-44-40.5T127 27zm-14 403V231q-4 4-49 45t-46 42q30 73 95 112zM351 36v272h98q13-35 13-76q0-60-29.5-112.5T351 36z"/>`),
		g.Group(children),
	)
}