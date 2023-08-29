package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PizzaBoyGba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.677 24.152l1.851 1.851m0-1.851l-1.851 1.851"/><circle cx="20.635" cy="23.357" r="3.617" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.859" cy="23.808" r="3.979" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="25.449" cy="34.697" r="2.962" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.38 13.27c-.33 1.75-.91 4.02-1.65 6.48"/><ellipse cx="40.96" cy="28.808" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.537" ry="1.834" transform="rotate(-16.754 40.96 28.807)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.09 19.844c-3.32-3.19-6.99-5.48-8.98-6.24a2.79 2.79 0 0 0-.34-.12a2.542 2.542 0 0 1-1.31-3.34c.56-1.29 2.01-1.68 3.34-1.32c1.91.51 6.05 3.01 9.96 6.55c3.92 3.53 7.61 8.09 8.63 12.7m-4.86 1.46c-.55-1.89-1.52-3.7-2.71-5.39m-11.51 13.3c-4.09.96-8.4 1.6-12.4 1.86c2.02-1.93 4.54-7.33 6.63-13m9.26 10.2c2.04-.61 3.98-1.3 5.73-2.07c.72-.32.24 2.24 2.2 2.03c.99-.1 1.52-.79 1.66-2.66c.11-1.46 1.79-2.69 2.34-3.2"/><circle cx="24.752" cy="18.615" r=".75" fill="currentColor"/><circle cx="24.252" cy="29.308" r=".75" fill="currentColor"/><circle cx="18.071" cy="36.244" r=".75" fill="currentColor"/><circle cx="34.028" cy="30.71" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.95 25.77c-1.51 1.2-3.81 2.76-7 4.27c-.67.31-1.36-.03-1.81-.49c-.89-.89-4.29-4.67-4.29-4.67c-.39-.5-.32-1.18-.21-1.57c.18-.68 1.48-2.78 1.99-5.18l7.99-7.99c2.4-.51 4.5-1.81 5.18-1.99c.38-.1 1.07-.18 1.56.22c0 0 .36.32.89.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.46 21.64v.01l-2.84 2.83c-.43.43-1.12.43-1.54 0l-3.97-3.96c-.42-.43-.42-1.12 0-1.54l6.36-6.36c.42-.42 1.11-.42 1.54 0l3.97 3.97c.42.42.42 1.11 0 1.54l-2.02 2.01"/><circle cx="20.459" cy="13.138" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}