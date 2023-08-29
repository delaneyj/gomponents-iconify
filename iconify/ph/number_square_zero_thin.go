package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareZeroThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 76c-12.82 0-23.41 5.81-30.62 16.81C91.33 102 88 114.52 88 128s3.33 26 9.38 35.2c7.21 11 17.8 16.8 30.62 16.8s23.41-5.81 30.62-16.8c6-9.22 9.38-21.72 9.38-35.2s-3.33-26-9.38-35.19C151.41 81.81 140.82 76 128 76Zm0 96c-22.11 0-32-22.1-32-44s9.89-44 32-44s32 22.1 32 44s-9.89 44-32 44Zm80-136H48a12 12 0 0 0-12 12v160a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12Zm4 172a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V48a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}