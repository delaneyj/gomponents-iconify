package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brilliant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5v43M38.19 7.85c-4.063 9.8-15.784 29.84-24.878 34.807"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.302 33.479c-7.09-9.899-18.56-21.717-37.358-21.151"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.268 20.843C34.46 30.942 19.392 35.18 4.865 33.81"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.28 4.153C39.926 17.194 42.19 36.533 24 45.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.044 12.308c-9.237.261-12.109 1.432-18.044 3.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 40.766c5.647 0 13.505-1.38 19.302-7.287C43.706 20.843 39.807 11.898 24 4.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.511c4.939.318 6.87 1.246 12.128 4.172M9.81 7.85c4.063 9.8 15.784 29.84 24.878 34.807"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.698 33.479c7.09-9.899 18.56-21.717 37.358-21.151"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.732 20.843C13.54 30.942 28.608 35.18 43.135 33.81"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.72 4.153C8.074 17.194 5.81 36.533 24 45.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.955 12.308c9.238.261 12.11 1.432 18.045 3.778"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 40.766c-5.647 0-13.505-1.38-19.302-7.287C4.294 20.843 8.193 11.898 24 4.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.511c-4.939.318-6.87 1.246-12.128 4.172"/>`),
		g.Group(children),
	)
}