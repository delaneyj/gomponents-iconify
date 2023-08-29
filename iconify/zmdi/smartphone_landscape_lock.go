package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneLandscapeLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 43q17 0 29.5 12.5T469 85v214q0 17-12.5 29.5T427 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43h384zm-43 256V85H85v214h299zm-192-22q-9 0-15-6t-6-15v-64q0-9 6-15t15-6v-22q0-17 12.5-29.5t30-12.5t30 12.5T277 149v22q9 0 15.5 6t6.5 15v64q0 9-6.5 15t-15.5 6h-85zm17-128v22h51v-22q0-10-7.5-17.5t-18-7.5t-18 7.5T209 149z"/>`),
		g.Group(children),
	)
}