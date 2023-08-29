package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 32C0 14.3 14.3 0 32 0h72.9c27.5 0 52 17.6 60.7 43.8L257.7 320c30.1.5 56.8 14.9 74 37l202.1-67.4c16.8-5.6 34.9 3.5 40.5 20.2s-3.5 34.9-20.2 40.5L352 417.7c-.9 52.2-43.5 94.3-96 94.3c-53 0-96-43-96-96c0-30.8 14.5-58.2 37-75.8L104.9 64H32C14.3 64 0 49.7 0 32zm244.8 102.5c-5.5-16.8 3.7-34.9 20.5-40.3L311 79.4l19.8 60.9l60.9-19.8l-19.9-60.9l45.7-14.8c16.8-5.5 34.9 3.7 40.3 20.5l49.4 152.2c5.5 16.8-3.7 34.9-20.5 40.3l-152.2 49.4c-16.8 5.5-34.9-3.7-40.3-20.5l-49.4-152.2z"/>`),
		g.Group(children),
	)
}