package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 3.012V2.5a.5.5 0 0 0-1 0v.512C6.898 3.149 3.142 6.84 3 11.5a.5.5 0 0 0 .5.5h8v7.25a1.75 1.75 0 0 1-3.5 0a.5.5 0 0 0-1 0A2.753 2.753 0 0 0 9.75 22a2.753 2.753 0 0 0 2.75-2.75V12h8a.5.5 0 0 0 .5-.5c-.142-4.66-3.898-8.35-8.5-8.488zM8.05 11H4.019a7.81 7.81 0 0 1 5.985-6.79C8.905 5.686 8.229 8.318 8.05 11zm1.002 0C9.345 6.664 10.776 4 12 4s2.656 2.664 2.948 7H9.052zm6.898 0c-.179-2.682-.854-5.314-1.952-6.79A7.81 7.81 0 0 1 19.982 11H15.95z"/>`),
		g.Group(children),
	)
}