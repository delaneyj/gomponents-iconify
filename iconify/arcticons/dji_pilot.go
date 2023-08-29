package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DjiPilot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="12" cy="11.977" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.5" ry="6.513"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.061 10.914l1.775-1.778m-1.775 3.904l5.194 5.204"/><circle cx="12" cy="11.977" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.164 14.818l1.776-1.778"/><ellipse cx="36" cy="11.977" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.5" ry="6.513"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.061 13.04l1.775 1.778m-3.897-1.778l-5.194 5.204"/><circle cx="36" cy="11.977" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.164 9.136l1.776 1.778"/><ellipse cx="36" cy="36.023" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.5" ry="6.513"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.939 37.086l-1.775 1.778m1.775-3.904l-5.194-5.204"/><circle cx="36" cy="36.023" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.836 33.182l-1.776 1.779"/><ellipse cx="12" cy="36.023" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.5" ry="6.513"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.939 34.96l-1.775-1.778m3.897 1.778l5.194-5.204"/><circle cx="12" cy="36.023" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.836 38.864l-1.776-1.778m5.195-18.842h0a33.433 33.433 0 0 0 11.49 0h0l-.006.033a33.435 33.435 0 0 0 0 11.446l.006.033h0a33.433 33.433 0 0 0-11.49 0h0l.006-.033a33.435 33.435 0 0 0 0-11.446Z"/>`),
		g.Group(children),
	)
}