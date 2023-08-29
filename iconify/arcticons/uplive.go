package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uplive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.982 16.054a5.41 5.41 0 0 0-3.873-3.052c-2.736-.391-3.182 2.233-3.238 3.46s.726 4.579 1.06 7.09s-.837 3.07-3.293 3.07s-3.182-1.06-3.406-2.4s2.513-6.085 2.903-8.709s-.837-3.405-2.4-3.405c-2.345 0-5.304 2.345-6.364 6.42s-1.564 9.1.949 12.45a10.674 10.674 0 0 0 11.165 3.46a11.18 11.18 0 0 0 4.194-2.298m9.567-17.938c1.252-.323 7.676 3.322 7.592 5.722s-6.42 4.634-7.592 4.773s-2.177-2.987-1.842-5.862s.977-4.41 1.842-4.633Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.091c0-2.348-1.842-9.685-9.853-9.909s-10.05 6.28-10.719 8.905s-4.625 13.112 2.289 18.255a3.778 3.778 0 0 0 3.805-.093c1.675-.977 1.565-3.737 1.162-4.844c-.306-.841-1.58-1.782-1.571-2.47c.028-2.14 3.582-2.447 6.736-2.419s8.151-2.428 8.151-7.424Z"/>`),
		g.Group(children),
	)
}