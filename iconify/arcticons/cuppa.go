package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cuppa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.999 23.395C7.438 24.822 4.5 27.06 4.5 29.576c0 4.31 8.618 7.804 19.25 7.804S43 33.886 43 29.576c0-2.516-2.938-4.754-7.498-6.181"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 29.786c0 4.31 4.682 9.885 19.25 9.885S43 33.886 43 29.576"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.182 12.99c0 6.115 4.08 18.064 14.568 18.064s14.567-11.949 14.567-18.063"/><ellipse cx="23.75" cy="12.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="14.567" ry="4.662"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.652 14.386c-1.704-1.944-7.283-3.369-13.902-3.369s-12.198 1.425-13.903 3.37m26.3 7.739c6.04 1.08 7.353-2.357 7.353-4.943s-.992-4.192-3.276-4.192a2.481 2.481 0 0 0-1.938.944"/>`),
		g.Group(children),
	)
}