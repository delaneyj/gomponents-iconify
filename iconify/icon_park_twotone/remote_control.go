package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRemoteControl0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="26" height="40" x="11" y="4" rx="2"/><circle cx="24" cy="34" r="4"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M18 10h12v8H18z"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 24h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRemoteControl0)"/>`),
		g.Group(children),
	)
}