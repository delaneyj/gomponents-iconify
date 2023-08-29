package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wechat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20 32.17a21.33 21.33 0 0 1-5.81-.71a2.63 2.63 0 0 0-1.45.24c-1.35.76-2.64 1.62-4.17 2.58c.28-1.27.46-2.38.78-3.45a1.31 1.31 0 0 0-.6-1.74c-4.65-3.28-6.61-8.2-5.14-13.26c1.36-4.68 4.69-7.51 9.22-9a15.47 15.47 0 0 1 16.89 4.94a15.09 15.09 0 0 1 2.71 7.55m-18.13-3a1.79 1.79 0 0 0-1.72-1.78a1.73 1.73 0 0 0-1.81 1.68a1.71 1.71 0 0 0 1.7 1.78a1.76 1.76 0 0 0 1.83-1.68Zm9.3-1.78a1.79 1.79 0 0 0-1.75 1.76a1.75 1.75 0 0 0 1.8 1.7a1.73 1.73 0 1 0 0-3.46Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.83 42.05c-1.22-.54-2.35-1.36-3.55-1.49s-2.45.57-3.69.69a12.36 12.36 0 0 1-10-3.27C18.22 33 19 25.46 25.17 21.41c5.49-3.6 13.55-2.4 17.42 2.6a9.88 9.88 0 0 1-1.14 13.8c-1.19 1.06-1.62 1.94-.86 3.33a4 4 0 0 1 .24.91Zm-14-13.52a1.44 1.44 0 1 0 0-2.88a1.44 1.44 0 1 0 0 2.88Zm9-2.88A1.43 1.43 0 0 0 34.44 27a1.42 1.42 0 1 0 2.84.12a1.43 1.43 0 0 0-1.41-1.47Z"/>`),
		g.Group(children),
	)
}