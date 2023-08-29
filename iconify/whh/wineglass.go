package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wineglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M901 576q-99 99-227.5 122T449 659L278 830l88 88q18 18 18 43.5t-18 43.5t-43.5 18t-43.5-18L18 745Q0 727 0 701.5T18 658t43.5-18t43.5 18l88 88l171-171q-61-96-38.5-224.5T447 123Q551 20 687 0q59 58 115 114t120 120.5T1024 337q-20 136-123 239zM674 68q-95 23-169 97q-47 47-74 104t-31 115h545q8-22 11-34z"/>`),
		g.Group(children),
	)
}