package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRotateBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 52h-25.58L170 33.34A12 12 0 0 0 160 28H96a12 12 0 0 0-10 5.34L73.57 52H48a28 28 0 0 0-28 28v112a28 28 0 0 0 28 28h160a28 28 0 0 0 28-28V80a28 28 0 0 0-28-28Zm4 140a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V80a4 4 0 0 1 4-4h32a12 12 0 0 0 10-5.34L102.42 52h51.15L166 70.66A12 12 0 0 0 176 76h32a4 4 0 0 1 4 4Zm-32-92v20a12 12 0 0 1-12 12h-20a12 12 0 0 1-7.76-21.14a28.07 28.07 0 0 0-29 2.73A12 12 0 0 1 96.79 94.4a52.28 52.28 0 0 1 61.14-.91A12 12 0 0 1 180 100Zm-18.41 52.8a12 12 0 0 1-2.38 16.8a51.71 51.71 0 0 1-31.13 10.34a52.3 52.3 0 0 1-30-9.44A12 12 0 0 1 76 164v-20a12 12 0 0 1 12-12h20a12 12 0 0 1 7.76 21.14a28.07 28.07 0 0 0 29-2.73a12 12 0 0 1 16.83 2.39Z"/>`),
		g.Group(children),
	)
}