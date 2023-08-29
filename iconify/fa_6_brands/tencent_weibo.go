package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TencentWeibo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M72.3 495.8c1.4 19.9-27.6 22.2-29.7 2.9C31 368.8 73.7 259.2 144 185.5c-15.6-34 9.2-77.1 50.6-77.1c30.3 0 55.1 24.6 55.1 55.1c0 44-49.5 70.8-86.9 45.1c-65.7 71.3-101.4 169.8-90.5 287.2zM192 .1C66.1.1-12.3 134.3 43.7 242.4C52.4 259.8 79 246.9 70 229C23.7 136.4 91 29.8 192 29.8c75.4 0 136.9 61.4 136.9 136.9c0 90.8-86.9 153.9-167.7 133.1c-19.1-4.1-25.6 24.4-6.6 29.1c110.7 23.2 204-60 204-162.3C358.6 74.7 284 .1 192 .1z"/>`),
		g.Group(children),
	)
}