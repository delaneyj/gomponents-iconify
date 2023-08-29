package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1645 326q-10 236-332 651q-333 431-562 431q-142 0-240-263q-44-160-132-482q-72-262-157-262q-18 0-127 76l-77-98q24-21 108-96.5T256 167Q412 29 497 21q95-9 153 55.5T731 280q44 287 66 373q55 249 120 249q51 0 154-161q101-161 109-246q13-139-109-139q-57 0-121 26Q1070-11 1409 0q251 8 236 326z"/>`),
		g.Group(children),
	)
}