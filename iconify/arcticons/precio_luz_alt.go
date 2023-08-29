package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrecioLuzAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.348 30.377c-.306-1.042-.25-2.275-.51-3.484a11.412 11.412 0 0 0-1.948-4.577C1.736 12.728 9.153 4.607 16.271 4.5c10.24.052 13.39 9.083 10.228 15.446a44.954 44.954 0 0 0-2.712 4.784a15.692 15.692 0 0 0-1.344 5.71a33.866 33.866 0 0 1-11.1-.068"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.348 30.377c-1.022 1.153-.994 2.278-.14 3.165a89.7 89.7 0 0 0 11.414.23a2.37 2.37 0 0 0-.175-3.332"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.57 33.805a2.118 2.118 0 0 0 .164 2.473c1.252.43 9.15.438 10.283.171a2.134 2.134 0 0 0 .16-2.609"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.475 36.6a1.698 1.698 0 0 0 .183 2.147c.875.836 4.776.871 6.188.108c.78-.624.712-1.28.438-2.163m1.304 3.038h7.198m-2.812-14.893h13.6m-11.066-4.442h8.204c2.887-.068 3.034 1.686 3.046 2.171v8.781M33.57 17.798v3.977m1.332 12.694l1.13 3.145l2.851-.763"/><circle cx="26.348" cy="28.456" r=".75" fill="currentColor"/><circle cx="31.065" cy="28.44" r=".75" fill="currentColor"/><circle cx="35.797" cy="28.44" r=".75" fill="currentColor"/><circle cx="26.32" cy="34.652" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.521 19.385c.767-.02 1.444 1.05 1.972 1.05c.756-.04 1.563-1.066 2.335-1.018c.668-.016 1.575 1.05 2.203 1.018c.7 0 1.034-1.07 1.909-1.113"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.238 20.435l-2.426 9.405l-2.482-9.421"/><circle cx="36.246" cy="37.726" r="5.774" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}