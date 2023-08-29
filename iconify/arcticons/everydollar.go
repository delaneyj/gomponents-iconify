package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Everydollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 14.957l-14.696-4.004c-9.938.518-16.862 2.214-24.304 9.373l15.214 4.851C24 18.914 33.609 14.487 43.5 14.958Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 16.244c-10.472.283-18.558 7.662-23.362 14.868L7.185 26.355c.707-2.677 2.846-4.265 2.846-4.265"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.123 17.688c-6.547.283-19.688 12.624-22.42 19.36c0 0-6.223-2.88-10.645-4.334c-.016-3.251 1.41-4.786 1.41-4.786"/><ellipse cx="22.87" cy="14.674" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.899" ry="1.413"/>`),
		g.Group(children),
	)
}