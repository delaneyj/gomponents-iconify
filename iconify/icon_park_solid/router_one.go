package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouterOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRouterOne0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M10 24L4 38h40l-6-14H10Z"/><path fill="#fff" fill-rule="evenodd" d="M10 4v20V4Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 4v20"/><path fill="#fff" fill-rule="evenodd" d="M38 4v20V4Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 4v20"/><path fill="#fff" fill-rule="evenodd" d="M24 4v20V4Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 4v20M4 38v6h40v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRouterOne0)"/>`),
		g.Group(children),
	)
}