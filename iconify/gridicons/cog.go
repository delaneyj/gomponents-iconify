package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 12c0-.568-.061-1.122-.174-1.656l1.834-1.612l-2-3.464l-2.322.786a7.97 7.97 0 0 0-2.859-1.657L14 2h-4l-.479 2.396a7.97 7.97 0 0 0-2.859 1.657L4.34 5.268l-2 3.464l1.834 1.612a7.993 7.993 0 0 0 0 3.312L2.34 15.268l2 3.464l2.322-.786a7.97 7.97 0 0 0 2.859 1.657L10 22h4l.479-2.396a7.978 7.978 0 0 0 2.859-1.657l2.322.786l2-3.464l-1.834-1.612A8.01 8.01 0 0 0 20 12zm-8 4a4 4 0 1 1 0-8a4 4 0 0 1 0 8z"/>`),
		g.Group(children),
	)
}