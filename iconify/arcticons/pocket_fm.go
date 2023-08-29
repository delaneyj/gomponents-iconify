package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PocketFm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="17.138" height="11.996" x="12.821" y="18.671" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.055"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.498 21.51h6.117m2.677 0h2.294"/><circle cx="15.907" cy="28.421" r=".857" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.512" cy="28.421" r=".857" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.92" cy="27.246" r="1.714" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.334 18.67l7.756-4.97m6.084 3.424a2.293 2.293 0 0 1 2.282 2.038"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.306 43.5s.035-3.78.157-20.802s21.877-26.17 31.257-9.05c6.229 11.37-2.607 30.316-21.898 22.617L7.306 43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.174 15.407a4.184 4.184 0 0 1 4.14 3.755"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.174 13.653a6.29 6.29 0 0 1 6.012 5.51"/>`),
		g.Group(children),
	)
}