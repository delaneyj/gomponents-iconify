package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Htmlfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M831.405 1024h-767q-27 0-45.5-18.5T.405 960V65q0-27 18.5-45.5T64.405 1h448v352q0 13 9 22.5t23 9.5h351v575q0 27-18.5 45.5t-45.5 18.5zm-459-440q12-12 12-30t-12-30.5t-29-12.5t-29 13l-175 149q-12 13-12 30.5t12 30.5l175 149q12 13 29 13t29-12.5t12-30.5t-12-30l-146-119zm384 89l-174-149q-12-13-29-13t-29 12.5t-12 30.5t12 30l145 120l-145 119q-12 12-12 30t12 30.5t29 12.5t29-13l174-149q13-13 13-30.5t-13-30.5zm-180-673q26 0 44 18l256 257q19 19 19 46h-319V0z"/>`),
		g.Group(children),
	)
}