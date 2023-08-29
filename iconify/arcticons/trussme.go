package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trussme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="39.2" cy="8.8" r="3.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.82 5.5a3.3 3.3 0 1 1 0 6.6H5.5V5.5h3.32M35.9 39.19a3.3 3.3 0 1 1 6.6 0v3.32h-6.6v-3.32"/><circle cx="8.82" cy="39.2" r="3.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.12 39.2H35.9m3.3-3.31V12.1m-3.3-3.3H12.12"/><circle cx="24" cy="24" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.867 11.133l-7.575 7.575M11.154 36.867L18.72 29.3m-9.9-17.2v23.8"/><circle cx="39.2" cy="8.8" r=".75" fill="currentColor"/><circle cx="8.81" cy="8.8" r=".75" fill="currentColor"/><circle cx="39.2" cy="39.2" r=".75" fill="currentColor"/><circle cx="8.81" cy="39.2" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 13.8v-10m28.7 38.7h10"/><circle cx="39.2" cy="44.129" r="1.629" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="42.5" cy="44.129" r="1.629" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.9" cy="44.129" r="1.629" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}