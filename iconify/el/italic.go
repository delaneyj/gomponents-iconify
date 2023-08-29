package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M578.299 1114.803h152.713L714.133 1200H245.546l16.879-85.197h152.713L620.897 86.001H468.185L485.867 0h468.587l-17.683 86.001H784.059l-205.76 1028.802"/>`),
		g.Group(children),
	)
}