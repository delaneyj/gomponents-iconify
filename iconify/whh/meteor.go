package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meteor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M886 234q33-19 59.5-29.5T982 193l10-1q-90 92-171 157q53-13 104-20t75-8l24-1Q851 433 768 512q-103 98-176.5 192T498 851q-26 76-92.5 124.5T256 1024q-106 0-181-75T0 768q0-69 34-127.5t91-92.5q38-19 82.5-60t79.5-84t66-86.5t46-67t17-26.5q-1 37-11 96q30 0 72-33t80-80t72-94t55-80l20-33q0 22-1 55.5t-9 103t-22 97.5q43-57 83-121t58-99l19-36q4 7 8 19.5t5 46t-13 62.5q38-27 86-59t77-51l29-18q-63 105-138 234zM256 576q-80 0-136 56T64 768t56 136t136 56t136-56t56-136t-56-136t-136-56z"/>`),
		g.Group(children),
	)
}