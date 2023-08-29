package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.49 4.11a10.451 10.451 0 0 0-9.26-.74a1.163 1.163 0 0 0-.731 1.08v11.66c0 .78.789 1.314 1.514 1.024a8.558 8.558 0 0 1 7.582.608l1.135.68c.087.053.18.075.269.074a.503.503 0 0 0 .27-.073l1.134-.681a8.558 8.558 0 0 1 7.582-.608a1.104 1.104 0 0 0 1.514-1.025V4.45c0-.476-.29-.903-.731-1.08a10.451 10.451 0 0 0-9.259.742l-.51.306l-.51-.306Zm1.26 2.39a.75.75 0 0 0-1.5 0V16a.75.75 0 0 0 1.5 0V6.5Z" clip-rule="evenodd"/><path fill="currentColor" d="M2.725 19.042a6.5 6.5 0 0 1 6.55 0l1.087.634a3.25 3.25 0 0 0 3.276 0l1.087-.634a6.5 6.5 0 0 1 6.55 0l.103.06a.75.75 0 1 1-.756 1.296l-.103-.06a5 5 0 0 0-5.038 0l-1.088.634a4.75 4.75 0 0 1-4.786 0l-1.088-.634a5 5 0 0 0-5.038 0l-.103.06a.75.75 0 0 1-.756-1.296l.103-.06Z"/>`),
		g.Group(children),
	)
}