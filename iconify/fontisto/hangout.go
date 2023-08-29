package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hangout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.951 0C4.455 0 0 4.455 0 9.951s4.455 9.951 9.951 9.951h.586V24c5.737-2.693 9.366-8.781 9.366-14.049v-.008C19.903 4.452 15.451 0 9.96 0h-.008zm-.585 10.537l-1.17 2.343h-1.76l1.17-2.342h-1.76V7.024h3.512v3.512zm4.683 0L12.88 12.88h-1.76l1.17-2.342h-1.76V7.024h3.512v3.512z"/>`),
		g.Group(children),
	)
}