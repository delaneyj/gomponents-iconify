package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drmanhattan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 544q0 98-38 186.5t-102.5 153t-153 102.5t-186.5 38t-186.5-38t-153-102.5T38 730.5T0 544q0-174 110-306T387 73q8-32 34-52.5T480 0t59 20.5T573 73q167 33 277 165t110 306zM566 137q-11 25-34.5 40T480 192t-51.5-15t-34.5-40q-143 30-236.5 144T64 544q0 113 55.5 209T271 904.5T480 960t209-55.5T840.5 753T896 544q0-149-93.5-263T566 137zm-86 503q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}