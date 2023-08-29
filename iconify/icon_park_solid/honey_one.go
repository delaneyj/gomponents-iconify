package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoneyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHoneyOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="11" height="5" x="4.929" y="13.224" rx="2" transform="rotate(-46.025 4.929 13.224)"/><rect width="11" height="5" x="19.321" y="27.111" rx="2" transform="rotate(-46.025 19.321 27.11)"/><rect width="17" height="5" x="6.443" y="18.855" rx="2" transform="rotate(-46.025 6.443 18.855)"/><rect width="17" height="5" x="13.641" y="25.798" rx="2" transform="rotate(-46.025 13.64 25.798)"/><rect width="25" height="5" x="7.265" y="25.205" rx="2" transform="rotate(-46.025 7.265 25.205)"/><path fill="#fff" stroke-linecap="round" stroke-linejoin="round" d="m25.003 28.424l4.166-4.318l14.407 13.9l-4.166 4.318l-14.407-13.9Z"/><path stroke-linejoin="round" d="M21 40.25c0 2.071-1.79 3.75-4 3.75s-4-1.679-4-3.75c0-2.071 4-6.25 4-6.25s4 4.179 4 6.25Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHoneyOne0)"/>`),
		g.Group(children),
	)
}