package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcamalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M449 830V701q108-12 182-93.5T705 415t-74-193t-182-94V0q106 8 194 65t139 149.5T833 415t-51 200.5T643 765t-194 65zm128-543q0 40-47 68t-113.5 28T303 355t-47-68t47-68t113.5-28T530 219t47 68zM128 415q0 111 74 192.5T385 701v129q-107-8-195-65T51 615.5T0 415t51-200.5T190 65T385 0v128q-109 12-183 94t-74 193zm0 417q23 22 56.5 36.5T263 888t74.5 6.5T417 896t79-1.5t74.5-6.5t78.5-19.5t56-36.5q22 21 32.5 34t21 38.5T769 960q0 12-9 21t-28 15.5t-38.5 11.5t-51 8t-55 4.5t-60.5 2.5t-56.5 1h-108l-56.5-1l-60-2.5l-55-4.5l-51.5-8l-38.5-11.5L73 981l-9-21q0-30 10.5-55.5T96 866t32-34z"/>`),
		g.Group(children),
	)
}