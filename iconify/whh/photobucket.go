package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photobucket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704H724q-96 64-212 64t-212-64H64q-27 0-45.5-18.5T0 640V128q0-26 18.5-45T64 64h64V32q0-13 9.5-22.5T160 0h128q13 0 22.5 9.5T320 32v21Q410 0 512 0q116 0 212 64h236q27 0 45.5 19t18.5 45v512q0 27-18.5 45.5T960 704zM64 640h162q-98-110-98-256t98-256H64v512zM512 64q-87 0-160.5 43T235 223.5T192 384t43 160.5T351.5 661T512 704t160.5-43T789 544.5T832 384t-43-160.5T672.5 107T512 64zm448 64H798q98 110 98 256t-98 256h162V128zM704 384q0-79-56-135.5T512 192v-64q106 0 181 75t75 181h-64z"/>`),
		g.Group(children),
	)
}