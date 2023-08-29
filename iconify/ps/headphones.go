package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M425 173q-7-65-55.5-108.5T256 21T142.5 64.5T87 173q-38 6-62.5 32.5T0 267q0 40 31 68t76 28h21V192q0-52 38-90t90-38t90 38t38 90v171h21q45 0 76-28t31-68q0-35-24.5-61.5T425 173zM85 318q-19-5-30.5-19.5T43 267q0-37 42-52v103zm342 0V218q19 4 30.5 18.5T469 269q0 34-42 49z"/>`),
		g.Group(children),
	)
}