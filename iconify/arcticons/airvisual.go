package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airvisual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.68 5.33A16.09 16.09 0 0 1 19 6.64a16.843 16.843 0 0 1 1.44 1.12m.45-3.26a16.088 16.088 0 0 1 2.29 1.31c.54.38 1 .76 1.44 1.12"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.15 5.17c.71-.06 1.33-.06 1.85-.06a15.23 15.23 0 0 1 7 1.41a12.8 12.8 0 0 1 5.86 5.26c1.59 2.89 1.45 5.63 1.29 8.9a30.402 30.402 0 0 1-1.51 7.81c-1.22 3.93-2.55 8.22-6.39 11.81c-.33.32-.6.56-.67.62c-1.09.92-3 2.58-5.58 2.58s-4.47-1.65-5.58-2.58c-.06-.05-.34-.3-.67-.62c-3.84-3.59-5.17-7.88-6.39-11.81a30.402 30.402 0 0 1-1.51-7.81c-.16-3.27-.3-6 1.29-8.9A12.8 12.8 0 0 1 17 6.52c.36-.17.71-.33 1.06-.46"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.82 32.77a2.862 2.862 0 0 0 .73-.08c2.29-.36 3.13-1.55 6.45-3.42a8.59 8.59 0 0 1 3.2-1.18a5.172 5.172 0 0 1 1.57 0A8.59 8.59 0 0 1 28 29.27c3.32 1.87 4.16 3.06 6.45 3.42c.32.05.58.07.74.08m-24.78-7.83a2.88 2.88 0 0 0 1.28-.42c1.24-.78 1.18-2.24 2-4.34a8.82 8.82 0 0 1 2.18-3.72a7.76 7.76 0 0 1 4.89-1.74c2.37-.11 3.05.76 5.39.8a10.5 10.5 0 0 0 4.73-1.12a8.171 8.171 0 0 0 1.08 2.78a8.44 8.44 0 0 0 2.38 2.52m-17.41 2.94h4.13"/><ellipse cx="19" cy="25.27" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.55" ry="1.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.06 22.64h-4.12"/><ellipse cx="29" cy="25.27" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.55" ry="1.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.41 24.94a1.71 1.71 0 0 0-1.18.54c-.34.43-.31 1-.08 2c.64 2.79 1 4.38 2.42 5a2.94 2.94 0 0 0 1.25.26m24.77-7.8a1.71 1.71 0 0 1 1.18.54c.34.43.31 1 .08 2c-.64 2.8-1 4.38-2.42 5a3.05 3.05 0 0 1-1.27.26"/>`),
		g.Group(children),
	)
}