package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.934 8.965A8.002 8.002 0 0 0 1 10c0 1.892.657 3.631 1.756 5.001C3.564 16.01 4 17.125 4 18.306V22h9l.001-3H15a2 2 0 0 0 2-2v-2.929l1.96-.84c.342-.146.372-.494.224-.727l-2.25-3.54ZM3 10a6 6 0 0 1 11.95-.779l.057.442l1.543 2.425l-1.55.664V17h-3.998L11 20H6v-1.694c0-1.639-.591-3.192-1.685-4.556A5.966 5.966 0 0 1 3 10Zm18.154 8.102l-1.665-1.11A8.959 8.959 0 0 0 21 12a8.958 8.958 0 0 0-1.51-4.993l1.664-1.11A10.948 10.948 0 0 1 23 12c0 2.258-.68 4.356-1.846 6.102Z"/>`),
		g.Group(children),
	)
}