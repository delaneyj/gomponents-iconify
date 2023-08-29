package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Datadotai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.12 1.744L.015 10.009L0 10.023l11.986 12.219l.014.015l11.986-12.22l.014-.014l-8.115-8.273l-.006-.006Zm1.207 1.02h5.326L11.99 5.41zm3.422 3.43l3.027-3.053L22.081 9.5h-6.054ZM8.211 3.14l3.04 3.072L7.999 9.5h-6.08Zm.62 6.977L12 6.876l3.169 3.242L12 19.842zm7.328.402h5.862l-8.793 9.005Zm-14.24 0h5.915l2.958 9.006Z"/>`),
		g.Group(children),
	)
}