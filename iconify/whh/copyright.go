package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm222 304q-6 21-23 34.5T672 480q-8 0-16-3t-13-6t-11.5-10t-9-10t-7-10.5t-5.5-8.5q-39-48-98-48q-53 0-90.5 37.5T384 512t37.5 90.5T512 640q59 0 98-48q2-2 9.5-13.5t12.5-17t16-11.5t24-6q22 0 39 13.5t23 34.5q1 4 2 11t-3.5 23.5T714 656q-37 50-90 81t-112 31q-106 0-181-75t-75-181t75-181t181-75q61 0 114.5 31.5T717 371q12 11 16 26.5t2 25.5z"/>`),
		g.Group(children),
	)
}