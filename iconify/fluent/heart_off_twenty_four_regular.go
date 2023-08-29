package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOffTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l1.855 1.856a5.375 5.375 0 0 0-.5 8.044l7.895 7.896a.75.75 0 0 0 1.06 0l3.744-3.742l4.445 4.447a.75.75 0 0 0 1.061-1.061L3.28 2.22Zm11.933 14.054L12 19.484L4.635 12.12a3.875 3.875 0 0 1 .512-5.912l10.066 10.066Zm4.155-4.153l-2.033 2.032l1.06 1.06l2.037-2.034a5.38 5.38 0 0 0-7.612-7.6l-.821.823l-.823-.823A5.36 5.36 0 0 0 7.19 4.008l1.89 1.891c.374.183.724.43 1.034.74l1.357 1.358a.75.75 0 0 0 1.074-.013L13.88 6.64a3.88 3.88 0 0 1 5.49 0a3.876 3.876 0 0 1-.002 5.481Z"/>`),
		g.Group(children),
	)
}