package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSwitchThemes0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="M23 4v23h20.993L44 4H23Z" clip-rule="evenodd"/><path d="M31.005 44H18.658c-1.702 0-3.742-.568-5.11-2.387c-.925-1.23-1.543-3.03-1.543-5.613V25"/><path stroke-linejoin="round" d="m4 33l8.005-8l8.009 8"/><path d="M30 19h7m-7-7h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSwitchThemes0)"/>`),
		g.Group(children),
	)
}