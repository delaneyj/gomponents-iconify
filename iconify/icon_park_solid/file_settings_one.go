package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSettingsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileSettingsOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12"/><circle cx="34" cy="36" r="5" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 28v3m0 10v3m5.828-14l-2.121 2.121M29.828 40l-2.121 2.121M28 30l2.121 2.121M38 40l2.121 2.121M26 36h3m10 0h3M30 4v10h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileSettingsOne0)"/>`),
		g.Group(children),
	)
}