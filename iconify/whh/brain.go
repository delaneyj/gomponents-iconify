package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 512q-12 0-47-16t-76-32t-69-16q-26 0-58 16.5t-64 40t-66.5 47t-77.5 40t-86 16.5q-52 0-90-25.5T256 512q0-33 21.5-57.5t52-37.5t72.5-20.5t75-9.5t67-2v-1q-13 0-22.5-9.5T512 352t9.5-22.5T544 320q-182 0-273 63q-3 2-22 15.5T225 415t-22.5 12t-30 12t-33 6t-43.5 3q-45 0-70.5-38T0 320q0-70 44-131T160 87.5t164-64T512 0q139 0 257 51.5t186.5 140T1024 384q0 45-26 86.5T928 512zM462 640l114-32q0-25 65-60.5t95-35.5q66 0 113 33t47 63q0 25-65 60.5T736 704q-22 0-27 48t-5 208h-64q0-50-59-141T472 704q-12-6-19.5-21.5t-5-29T462 640z"/>`),
		g.Group(children),
	)
}