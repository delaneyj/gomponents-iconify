package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Racquet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1007.5 1007.5q-17.5 17.5-42 17.5t-42.5-17L754 839q-19-20-16-50l-37-37q-13-14-23-22t-29-16.5t-40-8.5H431q-79 4-159-28.5T128 579Q54 506 22 413T5.5 232t77-149.5T232 5.5T413 22t166 106q65 64 97.5 144T705 431v178q0 46 47 92l37 37q30-3 50 16l168 169q18 18 18 42.5t-17.5 42zM538 169Q446 77 327 66.5T132 132T66.5 327T169 538t211 102.5T575 575t65.5-195T538 169z"/>`),
		g.Group(children),
	)
}