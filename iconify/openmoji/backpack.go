package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#D22F27" d="M56 31.083c0-11.045-8.954-20-20-20s-20 8.955-20 20v33.875h40V31.083z"/><path fill="#EA5A47" d="M51.647 52.965v7.621H20.354v-7.492v1.927v-24.416c0-8.641 7.005-15.647 15.646-15.647s15.646 7.006 15.646 15.647V55.02"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M56 31.083c0-11.045-8.954-20-20-20s-20 8.955-20 20v33.875h40V31.083z"/><path d="M51.647 52.965v7.621H20.354v-7.492v1.927v-24.416c0-8.641 7.005-15.647 15.646-15.647s15.646 7.006 15.646 15.647V55.02m-27.742-5.565h22.994M31 9a5 5 0 0 1 10 0"/></g>`),
		g.Group(children),
	)
}