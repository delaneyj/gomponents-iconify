package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M585 114h68c46 0 84 37 84 84v364c0 46-38 84-84 84H84c-46 0-84-38-84-84V198c0-47 38-84 84-84h69c17-27 48-66 76-66h279c28 0 60 39 77 66zM366 540c103 0 187-84 187-187c0-104-84-188-187-188c-104 0-189 84-189 188c0 103 85 187 189 187zm3-300c62 0 112 51 112 113c0 61-50 112-112 112s-113-51-113-112c0-62 51-113 113-113z"/>`),
		g.Group(children),
	)
}