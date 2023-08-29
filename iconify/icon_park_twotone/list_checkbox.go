package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTListCheckbox0"><g fill="none"><path fill="#555" fill-rule="evenodd" d="M20 24h24h-24Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 24h24"/><path fill="#555" fill-rule="evenodd" d="M20 38h24h-24Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 38h24"/><path fill="#555" fill-rule="evenodd" d="M20 10h24h-24Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 10h24"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 34h8v8H4zm0-14h8v8H4zM4 6h8v8H4z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTListCheckbox0)"/>`),
		g.Group(children),
	)
}