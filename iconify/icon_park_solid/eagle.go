package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eagle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEagle0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 23c-4.04-7.043 3.624-11.136 8-12c14.541-12.844 26.485-.287 28 8c1.514 8.287 1.158 14.893 2 18c-6.463-8.7-10.877-7.159-12-5c-2.02 4.144-5.314 4.252-7 3c-4.04-3.314-10.476 3.202-13 7c4.847-8.7 5.505-14.273 5-16c-2.02-8.286-8.307-5.416-11-3Z"/><circle cx="23" cy="16" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEagle0)"/>`),
		g.Group(children),
	)
}