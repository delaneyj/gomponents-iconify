package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hidedroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.58 21.628l-2.521-2.44a2.494 2.494 0 1 0-3.544 3.51l.337.332m.002 4.991a15.13 15.13 0 0 0-1.098 5.632V40.9a1.602 1.602 0 0 0 1.604 1.6h27.307a1.601 1.601 0 0 0 1.565-1.6v-7.247a15.103 15.103 0 0 0-1.062-5.612m-.46-4.576l.81-.798a2.492 2.492 0 1 0-3.537-3.51l-2.542 2.645"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 25.663h37s-19.205-10.502-37 0Zm7.52-9.386h21.96L31.539 5.5L21.185 7.105L16.182 5.91Zm9.168 16.11a2.53 2.53 0 0 1 3.625 0"/><circle cx="17.536" cy="32.387" r="4.652" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.464" cy="32.387" r="4.652" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}