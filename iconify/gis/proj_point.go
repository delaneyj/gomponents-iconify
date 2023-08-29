package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="12.55" cy="12.55" r="9.05" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="7"/><circle cx="61.943" cy="61.436" r="7.5" fill="currentColor"/><path fill="currentColor" d="m53.119 27.743l-8.267 8.266L31.69 22.847l-8.57 8.573l13.16 13.16l-8.266 8.266h25.105z" color="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="M46.227 100c.53-3.205 1.219-6.374 2.112-9.47l-4.804-1.386c-.954 3.31-1.683 6.669-2.24 10.04zm3.603-14.068a63.935 63.935 0 0 1 3.93-8.775l-4.393-2.386a68.917 68.917 0 0 0-4.24 9.46zm6.382-12.9c1.753-2.694 3.69-5.249 5.84-7.578l-3.673-3.392c-2.372 2.569-4.474 5.35-6.357 8.241zm9.234-10.941a57.95 57.95 0 0 1 7.609-5.813l-2.703-4.208a62.944 62.944 0 0 0-8.266 6.317zm11.752-8.262a71.265 71.265 0 0 1 8.786-4l-1.766-4.677a76.242 76.242 0 0 0-9.401 4.281zm13.372-5.572A87.926 87.926 0 0 1 100 45.91L99.057 41a92.938 92.938 0 0 0-9.967 2.48z" color="currentColor"/>`),
		g.Group(children),
	)
}