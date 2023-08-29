package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organicmaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.276 15.526c1.132.914 2.742 2.307 5.006 2.307c5.832 0 9.445-7.704 11.447-7.704a14.886 14.886 0 0 0-8.23-2.481c-4.898 0-8.232 2.133-8.223 7.878Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.585 27.67c-2.206.643-4.405.877-5.134-.522C5.092 24.535 8.896 19 14.276 15.526a39.416 39.416 0 0 1 11.36-4.875m7.417 11.677c-1.132-.915-2.742-2.307-5.006-2.307c-5.832 0-9.445 7.704-11.447 7.704a14.885 14.885 0 0 0 8.231 2.48c4.897 0 8.232-2.132 8.222-7.877Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.325 10.31c2.34-.733 4.779-1.093 5.553.396c1.36 2.611-2.444 8.147-7.825 11.621a39.416 39.416 0 0 1-11.36 4.875"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.712 43.659s14.756-15.117 14.756-24.244a14.756 14.756 0 0 0-29.512 0c0 9.127 14.756 24.244 14.756 24.244Z"/>`),
		g.Group(children),
	)
}