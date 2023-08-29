package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaddyOld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M.02 438.84c-1.53-95.641 82.635-207.298 172.921-306l62.688 3.404C166.517 202.89 114.336 272.276 83.658 341.722c74.918-111.375 196.21-201.494 367.792-268.563L512 81.667c-122.684 35.679-256.145 120.001-332.173 208.79C249.385 230 328.171 183.854 414.04 150.758l59.018 7.332c-101.478 38.808-197.837 90.643-255.307 157.217c-30.748 35.618-54.627 85.198-55.047 123.535H.02z"/>`),
		g.Group(children),
	)
}