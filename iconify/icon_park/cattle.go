package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cattle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M11.4651 19.9231C10.6818 12.4814 16.5167 6 23.9995 6C31.4823 6 37.3173 12.4814 36.5339 19.9231L35.0464 34.055C34.4513 39.7083 29.6841 44 23.9995 44C18.315 44 13.5478 39.7083 12.9527 34.0549L11.4651 19.9231Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13.9995 31C17.4916 27.8746 25.4281 23.8297 33.9995 31"/><circle cx="19" cy="18" r="2" fill="#fff"/><circle cx="21" cy="34" r="2" fill="#fff"/><circle cx="29" cy="18" r="2" fill="#fff"/><circle cx="27" cy="34" r="2" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 7.91262C35.1667 5.91262 40.3923 2.30498 43 4.91262C45.6077 7.5203 43 9.9126 41 10.4126C38.5 11.0376 36.8 12.7128 36 13.9128"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.0962 7.91262C12.9295 5.91262 7.70391 2.30498 5.09619 4.91262C2.48848 7.5203 5.09619 9.9126 7.09619 10.4126C9.59619 11.0376 11.2962 12.7128 12.0962 13.9128"/><path stroke="#000" stroke-width="4" d="M12 25L13 34.5"/><path stroke="#000" stroke-width="4" d="M36 25L35 34.5"/></g>`),
		g.Group(children),
	)
}