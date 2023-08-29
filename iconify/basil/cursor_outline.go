package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.255 3.724c-.1-1.192 1.235-1.85 2.152-1.238l.296.198a90.863 90.863 0 0 1 12.638 10.18c.97.936.115 2.412-1.094 2.282l-5.266-.564a1.208 1.208 0 0 0-1.234.68l-2.129 4.697c-.514 1.135-2.207 1.03-2.544-.195a86.016 86.016 0 0 1-2.79-15.693l-.03-.347Zm1.517.142l.007.08a84.508 84.508 0 0 0 2.636 15.035l1.966-4.338c.48-1.06 1.604-1.676 2.76-1.552l4.806.515A89.38 89.38 0 0 0 6.87 3.93l-.098-.065Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}