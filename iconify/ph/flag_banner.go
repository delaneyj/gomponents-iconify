package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230.76 51.73A8 8 0 0 0 224 48H32a8 8 0 0 0-5.41 13.9l42.09 38.57l-42.56 46.1A8 8 0 0 0 32 160h133.62l-28.84 60.56a8 8 0 1 0 14.44 6.88l80-168a8 8 0 0 0-.46-7.71ZM173.23 144h-123l35.61-38.57a8 8 0 0 0-.47-11.33L52.57 64h158.76Z"/>`),
		g.Group(children),
	)
}