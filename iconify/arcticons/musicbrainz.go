package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicbrainz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="31.966" cy="12.004" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.952" cy="16.16" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.117" cy="23.342" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.117" cy="31.325" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.123" cy="35.749" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.054" cy="29.795" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.052" cy="18.987" r="1.413" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.221 17.843a7.138 7.138 0 0 0-3.721-3.182m1.146.495a8.017 8.017 0 0 0 3.606-1.933m6.118 4.225a11.737 11.737 0 0 0 .06 4.66m-.177 8.105s-2.834-3.453-3.458-3.952a6.216 6.216 0 0 0-3.741-1.164c-1.982.026-2.673-.873-3.554-1.561"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.966 25.366s2.782-2.074 4.737-2.024m-4.851 13.014c-3.215 1.828-3.073-.558-5.352-.312m3.138-4.904S28.2 34.811 26.5 36.043"/><ellipse cx="11.154" cy="32.475" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.055" ry="2.078" transform="rotate(-23.434 11.154 32.475)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.07 31.751V12.876l7.43-4.612m-7.43 16.147l7.43-4.612m-7.43-.531l7.43-4.612m5-11.173l16.519 9.537v21.96L26.5 44.517V3.483z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.5 44.518L4.981 34.98V13.02L21.5 3.482v41.036z"/>`),
		g.Group(children),
	)
}