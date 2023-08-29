package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aperture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M725 654q6-10 9-14l221-384q69 119 69 256q0 28-4 63zm38-192l-4-16L645 18q165 44 270 179zM654 299q-11-7-14-9L256 69Q375 0 512 0q28 0 63 4zm-192-38l-16 4L18 379q44-165 179-270zM299 370q-7 11-9 14L69 768Q0 649 0 512q0-28 4-63zm-38 192q3 13 4 16l114 428q-165-44-270-179zm109 163q11 7 14 9l384 221q-119 69-256 69q-28 0-63-4zm192 38q14-3 16-4l428-114q-44 165-179 270z"/>`),
		g.Group(children),
	)
}