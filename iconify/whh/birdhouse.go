package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Birdhouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004.5 555.5Q985 575 957.5 575T911 555l-27-27l-84 431q-3 21-23.5 42.5T736 1023H289q-20 0-40.5-21.5T225 959l-84-432l-28 28q-19 20-46.5 20t-47-19.5t-19.5-47T19 462L463 19q20-20 49-19q29-1 49 19l444 443q19 19 19 46.5t-19.5 47zM512.5 895q26.5 0 45-18.5t18.5-45t-18.5-45.5t-45-19t-45 19t-18.5 45.5t18.5 45t45 18.5zm.5-512q-53 0-90.5 37.5T385 511t37.5 90.5T513 639t90-37.5t37-90.5t-37-90.5t-90-37.5z"/>`),
		g.Group(children),
	)
}