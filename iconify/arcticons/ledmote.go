package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ledmote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="26.59" height="39" x="10.7" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.77"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 7.64a2.37 2.37 0 1 0 .05 0Zm8 0A2.35 2.35 0 1 0 26.35 10A2.35 2.35 0 0 0 24 7.64Zm7.87 0A2.35 2.35 0 1 0 34.26 10a2.34 2.34 0 0 0-2.39-2.35ZM16 17h.05a2.33 2.33 0 1 1-.05 0Zm8 0h0a2.35 2.35 0 1 1-2.35 2.35A2.35 2.35 0 0 1 24 17Zm7.87 0a2.35 2.35 0 1 0 2.39 2.35A2.35 2.35 0 0 0 31.87 17ZM16 26.32h.05a2.33 2.33 0 1 1-.05 0Zm8 0h0a2.35 2.35 0 1 1-2.35 2.35A2.35 2.35 0 0 1 24 26.32Zm7.87 0a2.35 2.35 0 1 0 2.39 2.35a2.35 2.35 0 0 0-2.39-2.35ZM16 35.66h.05a2.33 2.33 0 1 1-.05 0Zm8 0h0A2.35 2.35 0 1 1 21.65 38A2.35 2.35 0 0 1 24 35.66Zm7.87 0A2.35 2.35 0 1 0 34.26 38a2.35 2.35 0 0 0-2.39-2.35ZM13.74 14.58h20.87"/>`),
		g.Group(children),
	)
}