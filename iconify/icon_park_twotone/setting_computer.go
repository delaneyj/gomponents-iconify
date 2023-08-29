package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSettingComputer0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 6H9a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-5m-18 8v8m-10 0h20"/><circle cx="37" cy="13" r="3" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 20v-4m0-6V6m-6.062 10.5l3.464-2m5.196-3l3.464-2m-12.124 0l3.464 2m5.196 3l3.464 2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSettingComputer0)"/>`),
		g.Group(children),
	)
}