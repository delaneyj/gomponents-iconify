package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ebobirthday(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.33 15.16c-7.33.75-12.78 3.67-12.78 7.17c0 4.06 7.37 7.35 16.45 7.35s16.45-3.29 16.45-7.35c0-3.5-5.46-6.42-12.78-7.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.55 26.41c0 1.44.38 3.55 2 4.69c1.14.81 3.46-.17 5.2.38c1.42.46 2.59 2.66 4.3 2.9c1.53.21 3.29-1.57 5-1.57s3.43 1.78 5 1.57c1.71-.24 2.89-2.44 4.31-2.89c1.74-.56 4.06.42 5.2-.39c1.6-1.14 2-3.25 2-4.69"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.55 22.33v13.44c0 4.06 7.37 7.35 16.45 7.35s16.45-3.29 16.45-7.35V22.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.33 14.15v8.18C20.33 23.23 22 24 24 24s3.67-.74 3.67-1.64v-8.21"/><ellipse cx="24" cy="14.15" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.67" ry="1.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.15c0-.9-1.05-2.1-1.05-2.86m.05 0c-.26-.77-1.75-.7-1.75-2.77s1.41-2.7 1.41-4.4a4 4 0 0 1 2.18 3.7A3.9 3.9 0 0 1 23 11.29Z"/>`),
		g.Group(children),
	)
}