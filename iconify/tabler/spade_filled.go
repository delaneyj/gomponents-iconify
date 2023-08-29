package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpadeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.327 2.26a1395.065 1395.065 0 0 0-4.923 4.504c-.626.6-1.212 1.21-1.774 1.843a6.528 6.528 0 0 0-.314 8.245l.14.177c1.012 1.205 2.561 1.755 4.055 1.574l.246-.037l-.706 2.118A1 1 0 0 0 9 22h6l.118-.007a1 1 0 0 0 .83-1.31l-.688-2.065l.104.02c1.589.25 3.262-.387 4.32-1.785a6.527 6.527 0 0 0-.311-8.243a31.787 31.787 0 0 0-1.76-1.83l-4.938-4.518a1 1 0 0 0-1.348-.001z"/></g>`),
		g.Group(children),
	)
}