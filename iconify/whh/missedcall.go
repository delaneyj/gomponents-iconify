package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Missedcall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m887 224l119 119q18 18 18 43.5t-18 44t-43.5 18.5t-43.5-18L800 311L681 431q-18 18-43.5 18t-44-18.5t-18.5-44t19-43.5l119-119l-119-119q-19-18-19-43.5T593.5 18t44-18T681 18l119 119L919 18q18-18 43.5-18t43.5 18t18 43.5t-18 43.5zM648 805l87-87q15-14 52.5-12t76.5 12t63 22q56 26 82 65.5t4 62.5L906 974q-44 44-117.5 49.5t-159-23t-182-90.5T264 761T115 578T24.5 396t-23-159T51 120L157 14q21-21 55.5 4T286 99q18 27 29 96.5t-8 95.5l-87 87q18 65 95 160t172 172.5T648 805z"/>`),
		g.Group(children),
	)
}