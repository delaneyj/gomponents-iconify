package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampaignMonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31.781 5.693a1.22 1.22 0 0 0-1.699-.303L.218 26.306c.219.313.584.521.995.521h29.573c.672 0 1.213-.547 1.213-1.219V6.373a1.286 1.286 0 0 0-.219-.681zM1.917 5.391a1.22 1.22 0 0 0-1.918.984V25.64l14.219-11.651z"/>`),
		g.Group(children),
	)
}